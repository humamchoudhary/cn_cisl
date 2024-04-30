package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/handler"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.211.249:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,  GET")
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
		t_r.POST("/auth", handler.TeacherAuthHandler)
		t_r.POST("/logout", handler.TeacherLogoutHandler)
		t_r.POST("/create-reservation", handler.ReserveLabhandler)
		t_r.POST("/edit-reservation", handler.EditReservationHandler)
		t_r.POST("/delete-reservation", handler.DeleteReservationHandler)
	}

	a_r := router.Group("/admin")
	{
		a_r.POST("/login", handler.AdminLoginHandler)
		a_r.POST("/create-teacher", handler.AdminCreateTeacherHandler)
		a_r.POST("/logout", handler.AdminLogOutHandler)
		a_r.POST("/delete-reservation", handler.AdminDeleteReservation)
		a_r.GET("/auth", handler.AdminAuthHandler)

	}

	r_r := router.Group("/reservation")
	{
		r_r.GET("/getAll", handler.HandlerGetAll)

	}

	temp_r := router.Group("/temp")
	{
		temp_r.GET("/session", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": handler.GetSessionByKey(c, "admin")})

		})
	}
	router.Run(":8000")
}
