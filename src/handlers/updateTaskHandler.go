package handlers

import (
	"log"
	"net/http"
	"strconv"

	"example.com/Todo/src/database"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a POST Methhod
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	params := r.URL.Query()
	idStr := params.Get("id")

	// Convert the 'id' parameter to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	title := r.FormValue("edit_title")
	status := r.FormValue("edit_status")

	updateQuery := "UPDATE tasks SET title = $1, status = $2 WHERE taskid = $3"

	_, err = database.DB.Exec(updateQuery, title, status, id)
	if err != nil {
		log.Println("Error updating data in the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
