package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetreserveHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "reserve.tmpl", gin.H{
		"title": "Main website",
	})
}
