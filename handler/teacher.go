package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/humamchoudhary/cn_cisl/models"
	"golang.org/x/crypto/bcrypt"
)

func auth(teacher *models.Teacher, password string) error {
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password))
}

func TeacherLoginHandler(c *gin.Context) {
	type LoginRequest struct {
		Id       string `json:"id"`
		Password string `json:"password"`
	}
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	var teacher models.Teacher
	teacher.Id = loginRequest.Id
	teacher.GetTeacherByID()
	fmt.Println(teacher)
	if err := auth(&teacher, loginRequest.Password); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Incorrect password"})
		return
	}
	SetSessionKey(c, "teacher", teacher.Id)
	c.JSON(http.StatusOK, gin.H{"message": "Login Success", "uid": teacher.Id})
}

func ReserveLabhandler(c *gin.Context) {
	teacherID := GetSessionByKey(c, "teacher")
	if teacherID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	type ReservationInput struct {
		ID         string `json:"id"`
		ReserverId string `json:"reservationid" `
		Date       string `json:"date" `
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime" `
		Recursive  bool   `json:"recursive"`
	}
	var reservationinput ReservationInput
	if err := c.ShouldBindJSON(&reservationinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	r := models.Reservation(reservationinput)
	id, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	r.ID = id.String()
	r.ReserverId = teacherID.(string)
	fmt.Println(r)
	if err = r.CreateReservation(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": r.ID})
}

func TeacherLogoutHandler(c *gin.Context) {
	SetSessionKey(c, "teacher", nil)
}

func EditReservation(c *gin.Context) {
	teacherID := GetSessionByKey(c, "teacher")
	if teacherID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	type ReservationEditReq struct {
		ID         string `json:"id"`
		ReserverId string `json:"reservationid" `
		Date       string `json:"date" `
		StartTime  string `json:"startTime"`
		EndTime    string `json:"endTime" `
		Recursive  bool   `json:"recursive"`
	}

	var reservationinput ReservationEditReq
	if err := c.ShouldBindJSON(&reservationinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	r_in := models.Reservation(reservationinput)

	r_search := models.Reservation{
		ID: r_in.ID,
	}

	r_search.GetByID()
	if r_search.ReserverId == teacherID.(string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	r_search.Edit(r_in)
	c.JSON(http.StatusOK, gin.H{})

}
