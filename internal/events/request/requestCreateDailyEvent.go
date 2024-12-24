package request

import (
	"github.com/pelletier/go-toml/v2"
)

type CreateDailyEventRequest struct {
	Title        string         `json:"title" binding:"required"`         // Obligatorio
	EmailCreator string         `json:"email_creator" binding:"required"` // Obligatorio
	Start        toml.LocalDate `json:"start" binding:"required"`         // Obligatorio
	End          toml.LocalDate `json:"end" binding:"required"`           // Obligatorio
	Always       bool           `json:"always" binding:"required"`
	Days         []int          `json:"days"`
	InitTime     toml.LocalTime `json:"init_time" binding:"required"`
	EndTime      toml.LocalTime `json:"end_time" binding:"required"`
	Location     string         `json:"location"`     // Opcional
	Description  string         `json:"description"`  // Opcional
	MeetLink     string         `json:"meet_link"`    // Opcional
	Participants []string       `json:"participants"` // Opcional

}
