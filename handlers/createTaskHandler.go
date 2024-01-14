package handlers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/test/database"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a POST Methhod
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
	_, err = database.DB.Exec(insertQuery, title, status)
	if err != nil {
		log.Println("Error inserting data into the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data successfully submitted and stored in the database")
}
