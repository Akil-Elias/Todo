package handlers

import (
	"fmt"
	"log"
	"net/http"

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
	titleStr := params.Get("title")
	fmt.Printf("the title is %s", titleStr)

	deleteQuery := "DELETE FROM todos WHERE title = $1"
	_, err := database.DB.Exec(deleteQuery, titleStr)
	if err != nil {
		log.Println("Error deleting data from the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data successfully deleted from the database")
}
