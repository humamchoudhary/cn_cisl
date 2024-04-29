package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/handler"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Next()
	}
}

func main() {

	router := gin.Default()
	router.Use(CORSMiddleware())
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	t_r := router.Group("/teacher")
	{
		t_r.POST("/login", handler.TeacherLoginHandler)
		t_r.POST("/create-reservation", handler.ReserveLabhandler)
	}

	a_r := router.Group("/admin")
	{
		a_r.POST("/login", handler.AdminLoginHandler)
		a_r.POST("/create-teacher", handler.AdminCreateTeacherHandler)
	}

	r_r := router.Group("/reservation")
	{
		r_r.GET("/getAll", handler.HandlerGetAll)
	}

	router.Run(":8000")
}
