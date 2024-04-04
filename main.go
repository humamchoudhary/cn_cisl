package main

import (
	"github.com/humamchoudhary/cn_cisl/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "Main website",
	// 	})
	// })

	router.GET("/reserve", handler.GetreserveHandler)
	router.GET("/login", handler.GetLoginHandler)
	router.Run(":80")
}
