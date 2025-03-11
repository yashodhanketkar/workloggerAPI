package users

import (
	"time"
	"worklogger/db"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Register(user User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	db := db.ConnectDB()
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users(name, password, created_at, updated_at) VALUES($1, $2, $3, $4)",
		user.Name,
		string(password),
		time.Now(),
		time.Now(),
	)

	return err
}

func Login(user User) error {
	var password string

	db := db.ConnectDB()
	defer db.Close()
	if err := db.QueryRow("SELECT password FROM users WHERE name = $1", user.Name).Scan(&password); err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	return err
}
