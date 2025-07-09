package controllers

import (
	"github.com/arbeeorlar/initializers"
	"github.com/arbeeorlar/models"
	"github.com/arbeeorlar/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpController(c *gin.Context) {

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.Student
	initializers.DB.Where("email = ?", student.Email).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Already exists"})
		return
	}

	var errHash error
	student.Password, errHash = utils.GeneratePasswordhash(student.Password)
	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errHash.Error()})
		return
	}
	initializers.DB.Create(&student)
	response := models.Response{
		Status:  "success",
		Message: "User Created",
		Data:    student,
	}
	c.JSON(http.StatusOK, response)
}
