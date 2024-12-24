package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/internal/events"
	"syncal/internal/events/request"
	"syncal/internal/events/service"
)

type HandlerEvent struct {
	repository []events.Event
}

func NewHandlerEvent(repo []events.Event) *HandlerEvent {
	return &HandlerEvent{repository: repo}
}

func (h *HandlerEvent) Create(c *gin.Context) {

	var req request.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rta, event := service.CreateEvent(req)

	h.repository = append(h.repository, event)
	c.IndentedJSON(http.StatusCreated, gin.H{"event": rta})
}
