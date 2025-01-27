package request

import (
	"time"
)

type UpdateEventRequest struct {
	Title       string    `json:"title"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	MeetLink    string    `json:"meet_link"`
}
