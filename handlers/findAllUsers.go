package handlers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})

}
