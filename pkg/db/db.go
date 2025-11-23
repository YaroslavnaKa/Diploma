package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
)

var db *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT "");
CREATE INDEX scheduler_date ON scheduler(date);`

func Init(dbF string) error {
	_, er := os.Stat(dbF)
	var install bool
	if os.IsNotExist(er) {
		install = true
	}
	var err error
	db, err = sql.Open("sqlite", dbF)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	if install {
		err = createTable(db)
		if err != nil {
			return err
		}
	}
	return nil
}
func createTable(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
