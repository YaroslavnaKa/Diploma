package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func AddTask(task *Task) (int64, error) {

	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func Tasks(search string, limit int) ([]Task, error) {
	tasks := make([]Task, 0)
	var rows *sql.Rows
	var err error

	if search != "" {
		d, errParse := time.Parse("02.01.2006", search)

		if errParse == nil {
			query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date LIMIT ?`
			rows, err = db.Query(query, d.Format("20060102"), limit)
		} else {

			query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`
			searchParam := "%" + search + "%"
			rows, err = db.Query(query, searchParam, searchParam, limit)
		}
	} else {

		query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?`
		rows, err = db.Query(query, limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Task
		e := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if e != nil {
			return nil, e
		}
		tasks = append(tasks, t)
	}

	if er := rows.Err(); er != nil {
		return nil, er
	}

	return tasks, nil

}
func GetTask(id string) (*Task, error) {
	e := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	q := db.QueryRow(e, id)
	var t Task
	err := q.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return nil, err
	}
	return &t, nil

}
func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}
func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`
	r, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	count, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for deleting task`)
	}
	return nil
}
