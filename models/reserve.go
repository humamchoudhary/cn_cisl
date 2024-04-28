package models

import (
	"fmt"

	"github.com/humamchoudhary/cn_cisl/database"
)

type Reservation struct {
	ID         int    `json:"id" db:"id,primarykey" primarykey:"true"`
	ReserverId string `json:"reservationname" db:"reserverId"`
	Date       string `json:"date" db:"date"`
	StartTime  string `json:"startTime" db:"start_time"`
	EndTime    string `json:"endTime" db:"end_time"`
	Recursive  bool   `json:"recursive" db:"recursive"`
}

func GetReservationByID(id int) (Reservation, error) {
	var reservations []Reservation
	err := database.Read("Reservation", map[string]interface{}{"id": id}, &reservations)
	if err != nil {
		return Reservation{}, err
	}
	if len(reservations) == 0 {
		return Reservation{}, fmt.Errorf("reservation with ID %d not found", id)
	}
	return reservations[0], nil
}

func (reservation Reservation) CreateReservation() error {

	err := database.Insert("Reservation", reservation)
	if err != nil {
		return err
	}
	return nil
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
		"start_time": startTime,
		"end_time":   endTime,
	}
	err := database.Read("Reservation", condition, &reservations)
	if err != nil {
		panic(err)
	}
	return reservations
}

func GetReservationsByReserverID(reserverID string) []Reservation {
	var reservations []Reservation
	condition := map[string]interface{}{
		"reserverId": reserverID,
	}
	err := database.Read("Reservation", condition, &reservations)
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

func (reservation *Reservation) EditReservation(newReservation Reservation) error {
	newReservationMap := map[string]interface{}{
		"reserverId": newReservation.ReserverId,
		"date":       newReservation.Date,
		"start_time": newReservation.StartTime,
		"end_time":   newReservation.EndTime,
		"recursive":  newReservation.Recursive,
	}

	condition := map[string]interface{}{
		"id": reservation.ID,
	}
	err := database.Update("Reservation", condition, newReservationMap)
	if err != nil {
		return err
	}
	return nil
}
