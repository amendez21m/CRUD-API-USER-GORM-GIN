package handlers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

type UpdateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	var input UpdateUserInput
	errInput := c.ShouldBindJSON(&input)

	err := models.DB.First(&user, id)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if errInput != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errInput.Error()})
		return
	}

	// Create book

	models.DB.Model(&user).Updates(models.User{Username: input.Email, Password: input.Password, Email: input.Email})

	c.JSON(http.StatusOK, gin.H{"data": user})
}
