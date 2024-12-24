package events

import (
	"github.com/pelletier/go-toml/v2"
	"syncal/internal/users"
)

// Setters Events
func (e *Event) SetTitle(title string) {
	e.title = title
}

func (e *Event) SetCreatedBy(createdBy users.User) {
	e.createdBy = createdBy
}

func (e *Event) SetLocation(location string) {
	e.location = location
}

func (e *Event) SetDescription(description string) {
	e.description = description
}

func (e *Event) SetMeetLink(meetLink string) {
	e.meetLink = meetLink
}

func (e *Event) SetStart(start toml.LocalDateTime) {
	e.start = start
}

func (e *Event) SetEnd(end toml.LocalDateTime) {
	e.end = end
}

func (e *Event) Start() toml.LocalDateTime {
	return e.start
}

func (e *Event) End() toml.LocalDateTime {
	return e.end
}

func (e *Event) CreatedBy() users.User {
	return e.createdBy
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) Location() string {
	return e.location
}

func (e *Event) Description() string {
	return e.description
}

func (e *Event) MeetLink() string {
	return e.meetLink
}
