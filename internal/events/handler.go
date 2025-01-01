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

	rta, _ := service.CreateEvent(req)

	c.IndentedJSON(http.StatusCreated, gin.H{"event": rta})
}
