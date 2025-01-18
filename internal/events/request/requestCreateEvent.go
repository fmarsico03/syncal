package request

import (
	"time"
)

type CreateEventRequest struct {
	Title        string    `json:"title" binding:"required"`         // Obligatorio
	EmailCreator string    `json:"email_creator" binding:"required"` // Obligatorio
	Start        time.Time `json:"start" binding:"required"`         // Obligatorio
	End          time.Time `json:"end" binding:"required"`           // Obligatorio
	Location     string    `json:"location"`                         // Opcional
	Description  string    `json:"description"`                      // Opcional
	MeetLink     string    `json:"meet_link"`                        // Opcional
	Participants []string  `json:"participants"`                     // Opcional
	Days         []int     `json:"days"`
	Weeks        []int     `json:"weeks"`
	Months       []int     `json:"months"`
	Always       bool      `json:"always"`
}
