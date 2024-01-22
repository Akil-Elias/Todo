package handlers

import (
	"log"
	"net/http"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a DELETE Methhod
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
