package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Main website",
	})
}

func GetSignupHandler(c *gin.Context) {
	
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Main website",
	})
}
