package handlers

import (
	"html/template"
	"log"
	"net/http"

	"example.com/Todo/database"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

const dataTemplate = `
{{range .}}
<tr>
	<td>
		<span> 
			{{.Title}} - {{.Status}}
		</span>
		<button hx-delete="http://localhost:8080/delete?id={{.ID}}" method="delete" hx-trigger="click" class="fa-solid fa-trash-can"></button>
	</td>
</tr>
{{end}}
`

var tmpl = template.Must(template.New("data").Parse(dataTemplate))

func GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s\n", r.Method, r.URL.Path)

	//Checks if the request is a GET Methhod
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	getAllQuery := "SELECT * FROM tasks"
	rows, err := database.DB.Query(getAllQuery)
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, "Error rendering HTML", http.StatusInternalServerError)
		return
	}
}
