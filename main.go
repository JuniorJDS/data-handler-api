package main

import (
	"log"

	"github.com/JuniorJDS/data-handler-api/api"
	"github.com/JuniorJDS/data-handler-api/api/app"
)

func main() {
	handler := api.HttpHandler()

	err := app.Start("5000", handler)
	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}
