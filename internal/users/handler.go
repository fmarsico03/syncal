package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/database"
	"syncal/internal/users/models"
	"syncal/internal/users/service"
)

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if service.EmailDuplicate(user.Mail) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Email is duplicated"})
		return
	}

	database.Database.Create(&user)

	c.IndentedJSON(http.StatusCreated, user.ID)
}
