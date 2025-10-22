package main

import (
	"log"
	"net/http"
	"os"

	"consent-service-extensions/pkg/api"
)

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create and configure router
	router := api.NewRouter()

	// Start server
	addr := ":" + port
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
