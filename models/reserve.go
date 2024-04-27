package models

import (
	"github.com/humamchoudhary/cn_cisl/database"
)

type Reservation struct {
	ID              int    `json:"id" db:"id,primarykey" primarykey:"true"`
	ReservationName string `json:"reservationname" db:"reservationname"`
	Date            string `json:"date" db:"date"`
	StartTime       string `json:"startTime" db:"start_time"`
	EndTime         string `json:"endTime" db:"end_time"`
	Recursive       bool   `json:"recursive" db:"recursive"`
}

func (reservation Reservation) CreateReservation() {

	err := database.Insert("Reservation", reservation)
	if err != nil {
		panic(err)
	}
}

func GetReservationsByRecurrence(recursive bool) []Reservation {
	var reservations []Reservation
	err := database.Read("Reservation", map[string]interface{}{"recursive": recursive}, &reservations)
	if err != nil {
		panic(err)
	}
	return reservations
}
func GetReservationsByTime(startTime, endTime string) []Reservation {
	var reservations []Reservation
	condition := map[string]interface{}{
		"startTime": startTime,
		"endTime":   endTime,
	}
	err := database.Read("Reservation", condition, &reservations)
	if err != nil {
		panic(err)
	}
	return reservations
}
func GetReservationsBySubject(teacherName string) []Reservation {
	var reservations []Reservation
	err := database.Read("Reservation", map[string]interface{}{"teacherName": teacherName}, &reservations)
	if err != nil {
		panic(err)
	}
	return reservations
}
func (reservation *Reservation) DeleteReservation() {
	condition := map[string]interface{}{
		"id": reservation.ID,
	}
	err := database.Delete("Reservation", condition)
	if err != nil {
		panic(err)
	}
}

// func (reservation *Reservation) EditReservation(newReservation Reservation) {
// 	newReservationMap := map[string]interface{}{
// 		"name":      newReservation.Name, // Change "Name" to "name"
// 		"date":      newReservation.Date,
// 		"startTime": newReservation.StartTime,
// 		"endTime":   newReservation.EndTime,
// 		"recursive": newReservation.Recursive,
// 	}

// 	condition := map[string]interface{}{
// 		"id": reservation.ID,
// 	}
// 	err := database.Update("Reservation", condition, newReservationMap)
// 	if err != nil {
// 		panic(err)
// 	}
// }
