package api

import (
	"diploma/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	srch := r.FormValue("search")
	tasks, err := db.Tasks(srch, 50)
	if err != nil {
		sendJson(w, ErrorResponse{Error: err.Error()})
		return
	}
	sendJson(w, TasksResp{Tasks: tasks})

}
