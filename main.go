package main

import (
	"fmt"

	"github.com/humamchoudhary/cn_cisl/database"
	"github.com/humamchoudhary/cn_cisl/models"
)

func main() {
	database.ConnectDatabase()
	// CREATING
	// teacher := models.Teacher{Id: 123, Name: "test", Department: "computer"}
	// models.CreateTeacher(teacher)

	// SELECTING
	var teacher []models.Teacher
	models.GetTeacherByID(123, &teacher)

	fmt.Println(teacher[0])

	// router := gin.Default()
	// t_r := router.Group("/teacher")
	// {
	// 	t_r.POST("/login", handler.HandlerTeacherLogin)
	// }
	// router.LoadHTMLGlob("templates/*")

	// // //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// // router.GET("/", func(c *gin.Context) {
	// // 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// // 		"title": "Main website",
	// // 	})
	// // })

	// router.GET("/reserve", handler.GetreserveHandler)
	// router.GET("/login", handler.GetLoginHandler)
	// router.Run(":80")
}
