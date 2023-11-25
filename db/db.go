package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "worklogger"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	if DB, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createDB("users.sql")
	createDB("projects.sql")
	createDB("tasks.sql")

	fmt.Println("Connected to the database")
}

func CloseDB() {
	DB.Close()
	fmt.Println("Connection closed")
}

func createDB(fileName string) {
	schemaTasks, err := os.ReadFile(filepath.Join("db", "init", fileName))
	if err != nil {
		log.Fatal(err)
	}

	if _, err = DB.Exec(string(schemaTasks)); err != nil {
		log.Fatal(err)
	}
}
