package request

import (
	"syncal/internal/utils/utilRepeatType"
	"time"
)

type CreateEventRequest struct {
	Title        string                `json:"title" binding:"required"`         // Obligatorio
	EmailCreator string                `json:"email_creator" binding:"required"` // Obligatorio
	Start        time.Time             `json:"start" binding:"required"`         // Obligatorio
	End          time.Time             `json:"end" binding:"required"`           // Obligatorio
	Location     string                `json:"location"`                         // Opcional
	Description  string                `json:"description"`                      // Opcional
	MeetLink     string                `json:"meet_link"`                        // Opcional
	Participants []string              `json:"participants"`                     // Opcional
	Repeat       utilRepeatType.Repeat `json:"repeat"`
}
