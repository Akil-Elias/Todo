package handlers

import (
	"html/template"
	"log"
	"net/http"

	"example.com/Todo/src/database"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

const fetchDataTemplate = `
{{range .}}
<tr>
	<td x-data="{ open: false }">
		<span id="task_card{{.ID}}"> 
			{{.Title}} - {{.Status}}
		</span>

		<button hx-delete="http://localhost:8080/delete?id={{.ID}}" hx-trigger="click" class="fa-solid fa-trash-can"></button>

		<span>
			<button @click="open = true" hx-swap="innerHTML" hx-target="#task_card{{.ID}}" class="fa-regular fa-pen-to-square"></button>
		</span>

		<span x-show="open">
			<form id="edit_form" hx-put="http://localhost:8080/put?id={{.ID}}" hx-vals="#edit_form">
				<label for="edit_title">Title:</label>
				<input type="text" name="edit_title" placeholder="Enter Title"><br>

				<label for="edit_status">Status:</label>
				<input type="text" name="edit_status" placeholder="Enter Status"><br>
				<button type="submit">Submit</button>
			</form>
		</span>
	</td>
</tr>
{{end}}
`

var fetch_tmpl = template.Must(template.New("data").Parse(fetchDataTemplate))

func FetchAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a GET Methhod
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fetchAllQuery := "SELECT * FROM tasks"
	rows, err := database.DB.Query(fetchAllQuery)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil { //WTF
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "text/html")
	err = fetch_tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, "Error rendering HTML", http.StatusInternalServerError)
		return
	}
}
