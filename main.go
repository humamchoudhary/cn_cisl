package main

import (
	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/handler"
)

func main() {
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
}
