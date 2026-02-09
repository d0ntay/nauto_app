package nautoapp_api

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Application struct {
	Config     config
	HTTPLogger *slog.Logger
	AppLogger  *slog.Logger
}

type config struct {
	addr string
}

func NewApplication(config config) (*Application, error) {

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String("time", a.Value.Time().UTC().Format(time.RFC3339))
			}
			return a
		},
	})

	httpLogger := slog.New(jsonHandler)
	appLogger := slog.New(jsonHandler)

	return &Application{
		Config:     config,
		HTTPLogger: httpLogger,
		AppLogger:  appLogger,
	}, nil
}

func NewConfig(addr string) (*config, error) {
	addr_parts := strings.Split(addr, ":")
	_, err := strconv.Atoi(addr_parts[1])
	if err != nil {
		return nil, err
	}
	if addr == "" {
		addr = "localhost:8080"
	}
	return &config{
		addr: addr,
	}, nil
}
