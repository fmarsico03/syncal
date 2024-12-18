package request

import (
	"time"
)

type CreateEventRequest struct {
	Title string `json:"title" binding:"required"` // Obligatorio
	//CreatedBy   users.User          `json:"created_by" binding:"required"` // Obligatorio
	Start       time.Time `json:"start" binding:"required"` // Obligatorio
	End         time.Time `json:"end" binding:"required"`   // Obligatorio
	Location    string    `json:"location"`                 // Opcional
	Description string    `json:"description"`              // Opcional
	MeetLink    string    `json:"meet_link"`                // Opcional
}
