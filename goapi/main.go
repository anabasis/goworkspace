package main

import (
	"log"
	"os"

	"github.com/anabasis/goapi/route"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set, defaulting to 3000")
	}

	app := route.Router()

	// Start the server
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

// This is a simple Fiber web server that responds with "Hello, World!" on the root
