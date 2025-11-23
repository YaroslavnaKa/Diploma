package main

import (
	"diploma/pkg/db"
	"diploma/pkg/server"
	"log"
	"os"
)

func main() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	webDir := "./web"
	e := db.Init(dbFile)
	if e != nil {
		log.Fatalf("Error conection to db: %v", e)
	}

	err := server.Run(port, webDir)
	if err != nil {
		log.Fatalf("Server failed  to start: %v", err)
	}
}
