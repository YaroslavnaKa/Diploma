package api

import (
	"diploma/pkg/db"
	"encoding/json"
	"net/http"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "Error decoding JSON input"})
		return
	}
	if task.ID == "" {
		sendJson(w, ErrorResponse{Error: "Не указан id"})
		return
	}
	if task.Title == "" {
		sendJson(w, ErrorResponse{Error: "Title is required"})
		return
	}
	er := checkDate(&task)
	if er != nil {
		sendJson(w, ErrorResponse{Error: er.Error()})
		return
	}
	e := db.UpdateTask(&task)
	if e != nil {
		sendJson(w, ErrorResponse{Error: "Task not found"})
		return
	}
	sendJson(w, map[string]interface{}{})

}
