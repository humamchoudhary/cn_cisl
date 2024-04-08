package models

import (
	"fmt"

	"github.com/humamchoudhary/cn_cisl/database"
)

type Teacher struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Department string `json:"dprt" db:"department"`
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
func GetTeacherByID(id int, teachers *[]Teacher) {

	err := database.Read("Teacher", map[string]interface{}{"id": 123}, teachers)
	if err != nil {
		panic(err)
	}
}
