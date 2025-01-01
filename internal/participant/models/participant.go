package models

import (
	"syncal/internal/events/models"
	models2 "syncal/internal/users/models"
)

type Participant struct {
	user   models2.User
	event  models.Event
	attend bool
}

func NewParticipant(user models2.User, event models.Event, attend bool) *Participant {
	return &Participant{user: user, event: event, attend: attend}
}
