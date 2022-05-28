package handlers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := models.DB.First(&user, id)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})

}
