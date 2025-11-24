package api

import (
	"diploma/pkg/db"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		sendJson(w, ErrorResponse{Error: "id is required"})
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "task not found"})
		return
	}
	if task.Repeat == "" {
		er := db.DeleteTask(id)
		if er != nil {
			sendJson(w, ErrorResponse{Error: er.Error()})
			return
		}

	} else {
		now := time.Now()
		repeat := task.Repeat
		d, e := NextDate(now, task.Date, repeat)
		if e != nil {
			sendJson(w, ErrorResponse{Error: e.Error()})
			return
		}
		task.Date = d
		errr := db.UpdateTask(task)
		if errr != nil {
			sendJson(w, ErrorResponse{Error: errr.Error()})
			return
		}
	}
	sendJson(w, map[string]interface{}{})
}
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	err := db.DeleteTask(id)
	if err != nil {
		sendJson(w, ErrorResponse{Error: err.Error()})
		return
	}
	sendJson(w, map[string]interface{}{})
}
