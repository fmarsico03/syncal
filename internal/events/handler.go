package events

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/internal/events/request"
	"syncal/internal/events/service"
)

func Create(c *gin.Context) {

	var req request.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rta, err := service.CreateEvent(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"event id": rta})
}

func Search(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

func SearchById(c *gin.Context) {

}
