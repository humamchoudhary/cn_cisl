package models

import (
	"errors"
	"time"

	"github.com/humamchoudhary/cn_cisl/database"
)

type Class struct {
	Id         int           `json:"id" db:"id,primarykey" primarykey:"true" autoincrement:"true"`
	StartTime  time.Time     `json:"startTime"`
	Duration   time.Duration `json:"duration"`
	Recurring  bool          `json:"-"`
	ReservedBy int           `json:"reservedBy"`
}

func validateTimeRange(t time.Time) bool {
	hour := t.Hour()
	return hour >= 8 && hour < 16
}

func validateDuration(d time.Duration) bool {
	hours := int(d.Hours())
	return hours >= 1 && hours <= 3
}

func (class *Class) ReserveClass() error {
	if !validateTimeRange(class.StartTime) {
		return errors.New("invalid Time Range")
	}

	if !validateDuration(class.Duration) {
		return errors.New("invalid Time Duration")

	}
	err := database.Insert("Class", class)
	if err != nil {
		return err
	}
	return nil

}

func GetAllClasses(retClasses interface{}) error {
	conditions := map[string]interface{}{}
	err := database.Read("Class", conditions, retClasses)
	if err != nil {
		return err
	}
	return nil
}

