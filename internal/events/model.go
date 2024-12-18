package events

import (
	"syncal/internal/users"
	"time"
)

type Event struct {
	title       string
	createdBy   users.User
	location    string
	description string
	meetLink    string
	start       time.Time
	end         time.Time
}

func (e *Event) End() time.Time {
	return e.end
}

func (e *Event) SetEnd(end time.Time) {
	e.end = end
}

func (e *Event) Start() time.Time {
	return e.start
}

func (e *Event) SetStart(start time.Time) {
	e.start = start
}

func NewEvent() *Event {
	return &Event{}
}

func NewEventComplete(title string, createdBy users.User, start, end time.Time) *Event {
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
