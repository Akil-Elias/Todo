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
	<td>
		<span id="task_card{{.ID}}"> 
			{{.Title}} - {{.Status}}
		</span>

		<button hx-delete="http://localhost:8080/delete?id={{.ID}}" hx-trigger="click" class="fa-solid fa-trash-can"></button>

		<span>
			<button hx-get="http://localhost:8080/edit_task" hx-swap="innerHTML" hx-target="#task_card{{.ID}}" class="fa-regular fa-pen-to-square"></button>
		</span>
	</td>
</tr>
{{end}}
`

var tmpl = template.Must(template.New("data").Parse(fetchDataTemplate))

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
	tmpl.Execute(w, tasks)
	if err != nil {
		http.Error(w, "Error rendering HTML", http.StatusInternalServerError)
		return
	}
}
