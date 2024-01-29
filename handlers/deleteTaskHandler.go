package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/Todo/database"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a DELETE Methhod
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
	fmt.Printf("%d", id)

	deleteQuery := "DELETE FROM tasks WHERE taskid = $1"
	_, err = database.DB.Exec(deleteQuery, id)
	if err != nil {
		log.Println("Error deleting data from the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
