package main

import (
	"github.com/humamchoudhary/cn_cisl/models"
)

func main() {
	// CREATING RESERVATION

	reservation := models.Reservation{
		ID:              2, // Provide a unique ID for the reservation
		ReservationName: "John Doe",
		Date:            "11/11/2011",
		StartTime:       "01:00 AM",
		EndTime:         "01:50 AM",
		Recursive:       true,
	}
	reservation.CreateReservation()

	// reservation := models.Reservation{ID: 1}
	// fmt.Println(reservation)

	// CREATING

	// t := models.Teacher{
	// 	Id:         5,
	// 	Name:       "John Doe",
	// 	Department: "Computer Science",
	// }
	// t.CreateTeacher()

	// // SELECTING by id
	// teacher := models.Teacher{Id: 4}
	// teacher.GetTeacherByID()
	// fmt.Println(teacher)

	// // Seletecting by name

	// // Single
	// teacher = models.Teacher{Name: "John Doe"}
	// teacher.GetTeacherByName()
	// fmt.Println(teacher)

	// // Multiple
	// var teachers []models.Teacher
	// teacher = models.Teacher{Name: "John Doe"}
	// err := teacher.GetTeacherByName(&teachers)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(teachers)

	// router := gin.Default()
	// t_r := router.Group("/teacher")
	// {
	// 	t_r.POST("/login", handler.HandlerTeacherLogin)
	// }
	// router.LoadHTMLGlob("templates/*")

	// router.GET("/reserve", handler.GetreserveHandler)
	// router.GET("/login", handler.GetLoginHandler)
	// router.Run(":80")
}
