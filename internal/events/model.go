package events

import (
	"github.com/gin-gonic/gin"
	"syncal/internal/users"
	"time"
)

type Event struct {
	title       string
	createdBy   users.User
	location    string
	description string
	meetLink    string
	frequency   []Recurrence
}

func NewEvent() *Event {
	return &Event{}
}

func NewEventComplete(title string, createdBy users.User, frequency ...Recurrence) *Event {
	return &Event{title: title, createdBy: createdBy, frequency: frequency}
}

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

func (e *Event) setFrequency(frequency ...Recurrence) {
	e.frequency = frequency
}

func (e *Event) SetDescription(description string) {
	e.description = description
}

func (e *Event) SetMeetLink(meetLink string) {
	e.meetLink = meetLink
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

func (e *Event) Frequency() []Recurrence {
	return e.frequency
}

type Recurrence struct {
	dayOfWeek  DayOfWeek
	startedAt  time.Time
	finishedAt time.Time
}

func (r *Recurrence) ToJSON() gin.H {
	return gin.H{
		"day_of_week": r.dayOfWeek,
		"started_at":  r.startedAt.Format(time.RFC3339),
		"finished_at": r.finishedAt.Format(time.RFC3339),
	}
}

func NewRecurrence(startedAt time.Time, finishedAt time.Time, dayOfWeek DayOfWeek) *Recurrence {
	return &Recurrence{startedAt: startedAt, finishedAt: finishedAt, dayOfWeek: dayOfWeek}
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

//Setters Recurrence

func (r *Recurrence) SetDayOfWeek(dayOfWeek DayOfWeek) {
	r.dayOfWeek = dayOfWeek
}

func (r *Recurrence) SetStartedAt(startedAt time.Time) {
	r.startedAt = startedAt
}

func (r *Recurrence) SetFinishedAt(finishedAt time.Time) {
	r.finishedAt = finishedAt
}

type FrecuencyRequest struct {
	DayOfWeek  int       `json:"day_of_week" binding:"required"`
	StartedAt  time.Time `json:"started_at" binding:"required"`
	FinishedAt time.Time `json:"finished_at" binding:"required"`
}

func NewFrecuencyRequest(dayOfWeek int, finishedAt time.Time, startedAt time.Time) *FrecuencyRequest {
	return &FrecuencyRequest{DayOfWeek: dayOfWeek, FinishedAt: finishedAt, StartedAt: startedAt}
}

func (f *FrecuencyRequest) GetDayOfWeek() int {
	return f.DayOfWeek
}

func (f *FrecuencyRequest) GetStartedAt() time.Time {
	return f.StartedAt
}

func (f *FrecuencyRequest) GetFinishedAt() time.Time {
	return f.FinishedAt
}

type CreateEventRequest struct {
	Title string `json:"title" binding:"required"` // Obligatorio
	//CreatedBy   users.User          `json:"created_by" binding:"required"` // Obligatorio
	Frequency   []FrecuencyRequest `json:"frequency" binding:"required"` // Obligatorio
	Location    string             `json:"location"`                     // Opcional
	Description string             `json:"description"`                  // Opcional
	MeetLink    string             `json:"meet_link"`                    // Opcional
}
