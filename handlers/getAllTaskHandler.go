package handlers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/Todo/database"
)

func GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a GET Methhod
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Insert data into the PostgreSQL database
	getAllQuery := "SELECT * FROM todos"
	_, err = database.DB.Exec(getAllQuery)
	if err != nil {
		log.Println("Error retrieving data from the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data successfully retrieved from the database")
}
