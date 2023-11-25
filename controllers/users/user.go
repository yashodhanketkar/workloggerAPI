package users

import (
	"time"
	"worklogger/db"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

func Register(user User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO users(name, password) VALUES($1, $2)", user.Name, string(password))
	return err
}

func Login(user User) error {
	var password string
	if err := db.DB.QueryRow("SELECT password FROM users WHERE name = $1", user.Name).Scan(&password); err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	return err
}
