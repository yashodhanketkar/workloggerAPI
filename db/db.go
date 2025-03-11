package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "records.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDB() {
	var err error
	if DB, err = sql.Open("sqlite3", "records.db"); err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	createTables("projects.sql")
	createTables("tasks.sql")
	createTables("users.sql")
}

func createTables(filename string) {
	path := filepath.Join("db", "init", filename)
	schemeTasks, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(string(schemeTasks))
	if err != nil {
		log.Fatal(err)
	}
}
