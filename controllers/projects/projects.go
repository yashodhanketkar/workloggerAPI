package projects

import (
	"log"
	"worklogger/db"
)

type Project struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func Create(project Project) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO projects(name, url) VALUES($1,$2)", project.Name, project.URL)
	return err
}

func Get(id int) (Project, error) {
	db := db.ConnectDB()
	defer db.Close()
	var project Project
	row := db.QueryRow("SELECT * FROM projects WHERE id = $1", id)
	err := row.Scan(
		&project.Id,
		&project.Name,
		&project.URL,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	return project, err
}

func ListAll() ([]Project, error) {
	db := db.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM projectS")
	if err != nil {
		log.Println("Error querying projects", err)
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project

		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.URL,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			log.Println("Error querying project", err)
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, err
}

func Update(id int, project Project) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec(
		"update projects set name = $1, url = $2 where id = $3",
		project.Name, project.URL, id,
	)
	return err
}

func Delete(id int) error {
	db := db.ConnectDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM projects WHERE id = $1", id)
	return err
}
