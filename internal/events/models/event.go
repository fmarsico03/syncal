package models

import (
	"gorm.io/gorm"
	"syncal/internal/users/models"
	"time"
)

type Event struct {
	gorm.Model
	Title       string
	CreatedByID uint        // Clave foránea explícita
	CreatedBy   models.User `gorm:"foreignKey:CreatedByID"`
	Location    string
	Description string
	MeetLink    string
	Start       time.Time // Cambiado de toml.LocalDateTime a time.Time
	End         time.Time // Cambiado de toml.LocalDateTime a time.Time
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
