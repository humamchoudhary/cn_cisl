package main

import (
	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/handler"
)

func main() {

	router := gin.Default()
	t_r := router.Group("/teacher")
	{
		t_r.POST("/login", handler.TeacherLoginHandler)
		t_r.POST("/create_reservation", handler.ReserveLabhandler)
	}

	a_r := router.Group("/admin")
	{
		a_r.POST("/login")
	}

	router.Run(":8000")
}
