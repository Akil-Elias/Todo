package main

import (
	"net/http"

	"example.com/Todo/src/database"
	"example.com/Todo/src/handlers"
	"github.com/rs/cors"
)

const port = ":8080"

func main() {
	// Connect to database
	database.Init()

	// Handler CORS
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Adjust this according to your needs
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Hx-Current-Url", "Hx-Request", "Hx-Target"},
		Debug:          true,
	})

	// Register the handler functions with CORS middleware
	http.Handle("/", cors.Handler(http.HandlerFunc(handlers.HomePageHandler)))
	http.Handle("/create", cors.Handler(http.HandlerFunc(handlers.CreateTaskHandler)))
	http.Handle("/fetchAll", cors.Handler(http.HandlerFunc(handlers.FetchAllTaskHandler)))
	http.Handle("/delete", cors.Handler(http.HandlerFunc(handlers.DeleteTaskHandler)))
	http.Handle("/put", cors.Handler(http.HandlerFunc(handlers.UpdateTaskHandler)))

	http.ListenAndServe(port, nil)
}
