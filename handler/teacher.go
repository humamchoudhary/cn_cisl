package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humamchoudhary/cn_cisl/models"
	"golang.org/x/crypto/bcrypt"
)

func auth(teacher *models.Teacher, password string) error {
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password))
}

func TeacherLoginHandler(c *gin.Context) {
	type LoginRequest struct {
		Id       int    `json:"id"`
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
	c.JSON(http.StatusOK, gin.H{"message": "Login Success", "uid": teacher.Id})
}

func TeacherSignUpHandler(c *gin.Context) {
	type SignUpRequest struct {
		Name       string `json:"name"`
		Id         int    `json:"id"`
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

	c.JSON(http.StatusCreated, gin.H{"message": "Teacher created successfully"})
}
func TeacherReserverTimehandler(c *gin.Context) {
	fmt.Print("reserve")
}

// t := models.Teacher{
// 		Id:         4,
// 		Name:       "John Doe",
// 		Department: "Computer Science",
// 	}
// 	t.CreateTeacher()

// 	// SELECTING by id
// 	teacher := models.Teacher{Id: 123}
// 	teacher.GetTeacherByID()
// 	fmt.Println(teacher)

// 	// Seletecting by name

// 	// Single
// 	teacher = models.Teacher{Name: "John Doe"}
// 	teacher.GetTeacherByName()
// 	fmt.Println(teacher)

// 	// Multiple
// 	var teachers []models.Teacher
// 	teacher = models.Teacher{Name: "John Doe"}
// 	err := teacher.GetTeacherByName(&teachers)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(teachers)
