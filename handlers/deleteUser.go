package handlers

import (
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := models.DB.Delete(&user, id)
	if err.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, "Error to Delete")

	}
	c.JSON(http.StatusOK, "Success")
}
