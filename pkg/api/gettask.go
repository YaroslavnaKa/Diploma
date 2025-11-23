package api

import (
	"diploma/pkg/db"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		sendJson(w, ErrorResponse{Error: "Не указан идентификатор"})
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "task not found"})
		return
	}
	sendJson(w, task)
}
