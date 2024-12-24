package events

import (
	"github.com/pelletier/go-toml/v2"
	"syncal/internal/users"
)

type Event struct {
	title       string
	createdBy   users.User
	location    string
	description string
	meetLink    string
	start       toml.LocalDateTime
	end         toml.LocalDateTime
}

func NewEvent() *Event {
	return &Event{}
}

func NewEventComplete(title string, createdBy users.User, start, end toml.LocalDateTime) *Event {
	return &Event{title: title, createdBy: createdBy, start: start, end: end}
}

type Recurrence struct {
	days   []DayOfWeek
	weeks  []int
	months []int
}

// Declarar el tipo de enum
type DayOfWeek int

// Definir los valores del enum utilizando iota
const (
	Sunday    DayOfWeek = iota // 0
	Monday                     // 1
	Tuesday                    // 2
	Wednesday                  // 3
	Thursday                   // 4
	Friday                     // 5
	Saturday                   // 6
)

type EventComplex struct {
	event      Event
	recurrence Recurrence
}

type EventDaily struct {
	event  Event
	always bool
	days   []DayOfWeek
}

type EventWeekly struct {
	event  Event
	always bool
	day    DayOfWeek
	week   []int
}

type EventMonthly struct {
	event  Event
	always bool
	day    DayOfWeek
	week   int
	month  []int
}
