package main

import (
	"net/http"

	"example.com/test/database"
	"example.com/test/handlers"
	"github.com/rs/cors"
)

const port = ":8080"

func main() {
	// Connect to database
	database.Init()

	// Handler CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Adjust this according to your needs
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Hx-Current-Url", "Hx-Request", "Hx-Target"},
		Debug:          true,
	})

	// Register the handler functions with CORS middleware
	http.Handle("/", c.Handler(http.HandlerFunc(handlers.HomePageHandler)))
	http.Handle("/submit", c.Handler(http.HandlerFunc(handlers.CreateTaskHandler)))
	http.ListenAndServe(port, nil)
}
