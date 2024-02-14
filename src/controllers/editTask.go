package controllers

import (
	"html/template"
	"log"
	"net/http"
)

const editDataForm = `
	<form>
		<label for="edit_title">Title:</label>
		<input type="text" name="edit_title" placeholder="Enter Title"><br>

		<label for="edit_status">Status:</label>
		<input type="text" name="edit_status" placeholder="Enter Status"><br>
	</form>
	<button hx-put="http://localhost:8080/put?id={{.ID}}" type="submit" hx-trigger="click">Submit</button>
`

var tmpl = template.Must(template.New("data").Parse(editDataForm))

func EditTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a GET Methhod
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, editDataForm)

}
