package main

import (
	"github.com/humamchoudhary/cn_cisl/models"
)

func main() {
	// CREATING RESERVATION

	reservation := models.Reservation{
		ID:         2, // Provide a unique ID for the reservation
		ReserverId: "John Doe",
		Date:       "11/11/2011",
		StartTime:  "01:00 AM",
		EndTime:    "01:50 AM",
		Recursive:  true,
	}
	reservation.CreateReservation()
	reservation.Edit(models.Reservation{
		ReserverId: "ali",
	})

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
	// teacher.Search()
	// fmt.Println(teacher)

	// // Seletecting by name

	// // Single
	// teacher = models.Teacher{Name: "John Doe"}
	// teacher.Search()
	// fmt.Println(teacher)

	// // Multiple
	// var teachers []models.Teacher
	// teacher = models.Teacher{Name: "John Doe"}
	// err := teacher.Search(&teachers)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(teachers)

	// router := gin.Default()
	// t_r := router.Group("/teacher")
	// {
	// 	t_r.POST("/login", handler.TeacherLoginHandler)
	// }

	// a_r := router.Group("/admin")
	// {
	// 	a_r.POST("/login")
	// }
	// router.Run(":8000")
}
