package participant

import (
	"syncal/internal/events"
	"syncal/internal/users"
)

type Participant struct {
	user   users.User
	event  events.Event
	attend bool
}

func NewParticipant(user users.User, event events.Event, attend bool) *Participant {
	return &Participant{user: user, event: event, attend: attend}
}
