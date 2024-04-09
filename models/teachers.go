package models

import (
	"github.com/humamchoudhary/cn_cisl/database"
)

type Teacher struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Department string `json:"dprt" db:"department"`
}

func (teacher Teacher) CreateTeacher() {
	err := database.Insert("Teacher", teacher)
	if err != nil {
		panic(err)
	}

}

func (teacher *Teacher) GetTeacherByID() {

	var teachers []Teacher
	err := database.Read("Teacher", map[string]interface{}{"id": teacher.Id}, &teachers)
	if err != nil {
		panic(err)
	}

	*teacher = teachers[0]

}

func (t *Teacher) GetTeacherByName(teachersOut ...interface{}) error {
	conditions := map[string]interface{}{
		"name": t.Name,
	}

	var teachers []Teacher
	if len(teachersOut) > 0 && teachersOut[0] != nil {
		err := database.Read("Teacher", conditions, teachersOut[0])
		if err != nil {
			return err
		}
	} else {
		err := database.Read("Teacher", conditions, &teachers)
		if err != nil {
			return err
		}
		*t = teachers[0]
	}

	return nil
}
