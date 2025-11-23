package api

import (
	"diploma/pkg/db"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type TaskIdResponse struct {
	Id int64 `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "Error decoding JSON input"})
		return

	}
	if task.Title == "" {
		sendJson(w, ErrorResponse{Error: "title is required"})
		return
	}
	err = checkDate(&task)
	if err != nil {
		sendJson(w, ErrorResponse{Error: err.Error()})
		return
	}
	id, err := db.AddTask(&task)
	if err != nil {
		sendJson(w, ErrorResponse{Error: "Error add task"})
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	successResp := TaskIdResponse{Id: id}
	json.NewEncoder(w).Encode(successResp)
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format("20060102")
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return errors.New("date format is wrong")
	}
	now = now.Truncate(24 * time.Hour)

	if t.Before(now) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format("20060102")

		} else {
			nextDate, e := NextDate(now, task.Date, task.Repeat)
			if e != nil {
				return e // Ошибка в правиле
			}
			task.Date = nextDate
		}
	} else {
		if task.Repeat != "" {
			_, er := NextDate(now, task.Date, task.Repeat)
			if er != nil {
				return er
			}
		}
	}

	return nil
}
