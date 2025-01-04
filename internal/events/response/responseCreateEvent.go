package response

import (
	"time"
)

type CreateEventResponse struct {
	Title        string    `json:"title" binding:"required"` // Obligatorio
	CreatedBy    string    `json:"created_by"`               // Obligatorio
	Location     string    `json:"location"`                 // Opcional
	Description  string    `json:"description"`              // Opcional
	MeetLink     string    `json:"meet_link"`                // Opcional
	Start        time.Time `json:"start" binding:"required"`
	End          time.Time `json:"end" binding:"required"`
	Participants []string  `json:"participants"`
}

func NewCreateEventResponse(createdBy string, description string, end time.Time, location string, meetLink string, participants []string, start time.Time, title string) *CreateEventResponse {
	return &CreateEventResponse{CreatedBy: createdBy, Description: description, End: end, Location: location, MeetLink: meetLink, Participants: participants, Start: start, Title: title}
}
