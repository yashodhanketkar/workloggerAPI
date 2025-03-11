package tasks

import (
	"log"
	"worklogger/db"
)

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Project     int    `json:"project"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func Create(task Task) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec(
		"INSERT INTO tasks(name, description, project) VALUES($1,$2,$3)",
		task.Name,
		task.Description,
		task.Project,
	)
	return err
}

func Get(id int) (Task, error) {
	var task Task
	db := db.ConnectDB()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM tasks WHERE id = $1", id)
	err := row.Scan(
		&task.Id,
		&task.Name,
		&task.Description,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.Project,
	)
	return task, err
}

func ListAll() ([]Task, error) {
	db := db.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Println("Error querying tasks", err)
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Description,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.Project,
		)
		if err != nil {
			log.Println("Error querying task", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}

func Update(id int, task Task) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec(
		"UPDATE tasks SET name = $1, description = $2, project = $3 WHERE id = $4",
		task.Name, task.Description, task.Project, id,
	)
	return err
}

func Delete(id int) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
