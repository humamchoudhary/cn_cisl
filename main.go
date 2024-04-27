package main

import (
<<<<<<< HEAD
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
=======
	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/handler"
)

func main() {
>>>>>>> 0901d8ffdb650d633053e3160264d612c3e17c41

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

<<<<<<< HEAD
	// router := gin.Default()
	// t_r := router.Group("/teacher")
	// {
	// 	t_r.POST("/login", handler.HandlerTeacherLogin)
	// }
	// router.LoadHTMLGlob("templates/*")

	// router.GET("/reserve", handler.GetreserveHandler)
	// router.GET("/login", handler.GetLoginHandler)
	// router.Run(":80")
=======
	router := gin.Default()
	t_r := router.Group("/teacher")
	{
		t_r.POST("/login", handler.TeacherLoginHandler)
	}

	c_r := router.Group("/class")
	{
		c_r.GET("/all")
	}

	router.Run(":80")
>>>>>>> 0901d8ffdb650d633053e3160264d612c3e17c41
}
