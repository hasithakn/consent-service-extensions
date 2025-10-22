package main

import (
	"log"
	"net/http"

	"consent-service-extensions/internal/config"
	"consent-service-extensions/pkg/api"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create and configure router
	router := api.NewRouter()

	// Start server
	addr := ":" + cfg.Port
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Log level: %s", cfg.LogLevel)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
