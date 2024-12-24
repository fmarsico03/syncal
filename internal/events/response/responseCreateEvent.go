package response

import (
	"github.com/pelletier/go-toml/v2"
)

type CreateEventResponse struct {
	Title        string             `json:"title" binding:"required"` // Obligatorio
	CreatedBy    string             `json:"created_by"`               // Obligatorio
	Location     string             `json:"location"`                 // Opcional
	Description  string             `json:"description"`              // Opcional
	MeetLink     string             `json:"meet_link"`                // Opcional
	Start        toml.LocalDateTime `json:"start" binding:"required"`
	End          toml.LocalDateTime `json:"end" binding:"required"`
	Participants []string           `json:"participants"`
}

func NewCreateEventResponse(createdBy string, description string, end toml.LocalDateTime, location string, meetLink string, participants []string, start toml.LocalDateTime, title string) *CreateEventResponse {
	return &CreateEventResponse{CreatedBy: createdBy, Description: description, End: end, Location: location, MeetLink: meetLink, Participants: participants, Start: start, Title: title}
}
