package models

import (
	"database/sql"
	"fmt"

	"github.com/humamchoudhary/cn_cisl/database"
)

type Teacher struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"dprt"`
}

func CreateTeacher(teacher Teacher) error {
	database.DB.Exec("CREATE TABLE IF NOT EXISTS Teacher (id INTEGER PRIMARY KEY, name TEXT, department TEXT)")

	stmt, err := database.DB.Prepare("INSERT INTO Teacher(id,name, department) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("Error while prepareing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(teacher.Id, teacher.Name, teacher.Department) // Execute statement with values
	if err != nil {
		// Handle error
		fmt.Println("Error inserting", err)
		return err

	}
	return nil
}
func GetTeacherByID(id int) (Teacher, error) {
	database.DB.Exec("CREATE TABLE IF NOT EXISTS Teacher (id INTEGER PRIMARY KEY, name TEXT, department TEXT)")
	stmt, err := database.DB.Prepare("SELECT id,name,department from Teacher WHERE id = ?")
	if err != nil {
		fmt.Println("Error while prepareing", err)
	}
	defer stmt.Close()
	teacher := Teacher{}
	sqlErr := stmt.QueryRow(id).Scan(&teacher.Id, &teacher.Name, &teacher.Department)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Teacher{}, nil
		}
		return Teacher{}, sqlErr
	}

	return teacher, sqlErr
}
