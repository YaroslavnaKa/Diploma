package api

import (
	"encoding/json"
	"net/http"
)

const DateFormat = "20060102"

func Init() {
	http.HandleFunc("/api/signin", SigninHandler)
	http.HandleFunc("/api/nextdate", NextDateHandler)
	http.HandleFunc("/api/task", Auth(taskHandler))
	http.HandleFunc("/api/tasks", Auth(tasksHandler))
	http.HandleFunc("/api/task/done", Auth(doneTaskHandler))
}
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)

	case http.MethodGet:
		getTaskHandler(w, r)
	case http.MethodPut:
		editTaskHandler(w, r)

	case http.MethodDelete:
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func sendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
	}
}
