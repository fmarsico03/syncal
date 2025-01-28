package models

import (
	"gorm.io/gorm"
	"syncal/internal/users/models"
	"syncal/internal/utils/utilRepeatType"
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
	Start       time.Time
	End         time.Time
	LinkedId    uint
	Type        utilRepeatType.RepeatType
}
