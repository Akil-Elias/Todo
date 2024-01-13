package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

var db *sql.DB

func init() {
	// Initialize the database connection
	connStr := "postgresql://Akil-Elias:k7CI6oLEZeAM@ep-calm-leaf-80663857.us-east-1.aws.neon.tech/todoDB?sslmode=require"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Adjust this according to your needs
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Hx-Current-Url", "Hx-Request", "Hx-Target"},
		Debug:          true,
	})

	// Register the handler function with CORS middleware
	http.Handle("/submit", c.Handler(http.HandlerFunc(submitHandler)))
	http.ListenAndServe(":8080", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	status := r.FormValue("status")

	// Insert data into the PostgreSQL database
	insertQuery := "INSERT INTO todos (title, status) VALUES ($1, $2)"
	_, err = db.Exec(insertQuery, title, status)
	if err != nil {
		log.Println("Error inserting data into the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data successfully submitted and stored in the database")
}
