package nautoapp_api

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (app *Application) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)

		app.HTTPLogger.Info("http request",
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.String("real_ip", r.RemoteAddr),
			slog.Int("status", wrapped.statusCode),
			slog.Duration("duration", duration),
			slog.String("user_agent", r.UserAgent()),
		)
	})
}
