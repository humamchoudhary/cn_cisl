package models

import (
	"fmt"

	"github.com/humamchoudhary/cn_cisl/database"
)

type Reservation struct {
	ID         string `json:"id" db:"id,primarykey" primarykey:"true"`
	ReserverId string `json:"reservationname" db:"reserverId"`
	Date       string `json:"date" db:"date"`
	StartTime  string `json:"startTime" db:"start_time"`
	EndTime    string `json:"endTime" db:"end_time"`
	Recursive  bool   `json:"recursive" db:"recursive"`
}

func (reservation Reservation) CreateReservation() error {

	err := database.Insert("Reservation", reservation)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reservation) GetByID() error {
	var reservations []Reservation
	err := database.Read("Reservation", map[string]interface{}{"id": r.ID}, &reservations)
	if err != nil {
		return err
	}
	if len(reservations) == 0 {
		return fmt.Errorf("reservation with ID %d not found", r.ID)
	}
	*r = reservations[0]
	return nil
}

func GetReservationsByRecurrence(recursive bool) ([]Reservation, error) {
	var reservations []Reservation
	err := database.Read("Reservation", map[string]interface{}{"recursive": recursive}, &reservations)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationsByTime(startTime, endTime string) ([]Reservation, error) {
	var reservations []Reservation
	condition := map[string]interface{}{
		"start_time": startTime,
		"end_time":   endTime,
	}
	err := database.Read("Reservation", condition, &reservations)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func GetReservationsByReserverID(reserverID string) ([]Reservation, error) {
	var reservations []Reservation
	condition := map[string]interface{}{
		"reserverId": reserverID,
	}
	err := database.Read("Reservation", condition, &reservations)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *Reservation) Delete() error {
	condition := map[string]interface{}{
		"id": r.ID,
	}
	err := database.Delete("Reservation", condition)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reservation) Edit(newReservation Reservation) error {
	newReservationMap := make(map[string]interface{})

	if newReservation.ReserverId != "" {
		newReservationMap["reserverId"] = newReservation.ReserverId
	}

	if newReservation.Date != "" {
		newReservationMap["date"] = newReservation.Date
	}

	if newReservation.StartTime != "" {
		newReservationMap["start_time"] = newReservation.StartTime
	}

	if newReservation.EndTime != "" {
		newReservationMap["end_time"] = newReservation.EndTime
	}

	if newReservation.Recursive != r.Recursive {
		newReservationMap["recursive"] = newReservation.Recursive
	}

	condition := map[string]interface{}{
		"id": r.ID,
	}

	err := database.Update("Reservation", newReservationMap, condition)
	if err != nil {
		return err
	}

	return nil
}

func GetAll() error {
	return nil
}
