package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/models"
	"golang.org/x/crypto/bcrypt"
)

func AdminLoginHandler(c *gin.Context) {

	type Admin struct {
		Username string `json:'username'`
		Password string `json:'password'`
	}

	var admin Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if admin.Username == "humamAdmin" && admin.Password == "Admin123" {

		SetSessionKey(c, "admin", admin.Username)
		c.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		SetSessionKey(c, "admin", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username/password"})
		return

	}

}

func AdminCreateTeacherHandler(c *gin.Context) {
	if admin := GetSessionByKey(c, "admin"); admin == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	type SignUpRequest struct {
		Name       string `json:"name"`
		Id         string `json:"id"`
		Password   string `json:"password"`
		Department string `json:"department"`
	}
	var signUpRequest SignUpRequest
	if err := c.ShouldBindJSON(&signUpRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newTeacher := models.Teacher{
		Department: signUpRequest.Department,
		Name:       signUpRequest.Name,
		Id:         signUpRequest.Id,
		Password:   string(hashedPassword),
	}
	err = newTeacher.CreateTeacher()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher", "e": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher created successfully"})

}

func AdminLogOutHandler(c *gin.Context) {
	if admin := GetSessionByKey(c, "admin"); admin == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	SetSessionKey(c, "admin", nil)
	c.JSON(http.StatusOK, gin.H{})

}
