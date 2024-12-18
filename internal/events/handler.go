package events

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/internal/events/request"
	response2 "syncal/internal/events/response"
	"syncal/internal/users"
)

type HandlerEvent struct {
	repository []Event
}

func NewHandlerEvent(repo []Event) *HandlerEvent {
	return &HandlerEvent{repository: repo}
}

func (h *HandlerEvent) Create(c *gin.Context) {
	franco := users.NewUser("Franco", "Marsico", "fmarsico03@gmail.com")

	var req request.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event := NewEventComplete(req.Title, *franco, req.Start, req.End)
	event.SetDescription(req.Description)
	event.SetLocation(req.Location)
	event.SetMeetLink(req.MeetLink)

	fmt.Printf("Titulo: %s, Descripcion: %s, Meet: %s\n", event.Title(), event.Description(), event.MeetLink())

	event.SetCreatedBy(*franco)
	h.repository = append(h.repository, *event)

	rta := response2.NewCreateEventResponse(
		event.createdBy.Name()+" "+event.createdBy.Lastname(),
		event.description,
		event.end,
		event.location,
		event.meetLink,
		event.start,
		event.title)
	c.IndentedJSON(http.StatusCreated, gin.H{"event": rta})
}
