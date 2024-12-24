package request

import (
	"github.com/pelletier/go-toml/v2"
)

type CreateEventRequest struct {
	Title        string             `json:"title" binding:"required"`         // Obligatorio
	EmailCreator string             `json:"email_creator" binding:"required"` // Obligatorio
	Start        toml.LocalDateTime `json:"start" binding:"required"`         // Obligatorio
	End          toml.LocalDateTime `json:"end" binding:"required"`           // Obligatorio
	Location     string             `json:"location"`                         // Opcional
	Description  string             `json:"description"`                      // Opcional
	MeetLink     string             `json:"meet_link"`                        // Opcional
	Participants []string           `json:"participants"`                     // Opcional
}
