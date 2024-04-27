package models

import (
	"github.com/humamchoudhary/cn_cisl/database"
)

type Teacher struct {
	Id         int    `json:"id" db:"id,primarykey" primarykey:"true" `
	Name       string `json:"name" db:"name"`
	Department string `json:"dprt" db:"department"`
	Password   string `json:"-" db:"password"`
}

func (teacher Teacher) CreateTeacher() error {
	err := database.Insert("Teacher", teacher)
	if err != nil {
		return err
	}
	return nil
}

func (teacher *Teacher) GetTeacherByID() error{

	var teachers []Teacher
	err := database.Read("Teacher", map[string]interface{}{"id": teacher.Id}, &teachers)
	if err != nil {
		return err
	}

	*teacher = teachers[0]
return nil
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