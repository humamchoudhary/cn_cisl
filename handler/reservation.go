package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/models"
)

func HandlerGetAll(c *gin.Context) {
	r, err := models.GetAllReservations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": r})

}
