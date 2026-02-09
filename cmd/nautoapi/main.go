package main

import (
	"log"

	"github.com/d0ntay/nautoapp/internal/nautoapp_api"
)

func main() {
	cfg, err := nautoapp_api.NewConfig(":8080")
	if err != nil {
		log.Fatal(err)
	}
	app, err := nautoapp_api.NewApplication(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	mux := app.MountRouter()
	log.Fatal(app.StartServer(mux))
}
