package models

import (
	"gorm.io/gorm"
	modelsEvent "syncal/internal/events/models"
	modelsUser "syncal/internal/users/models"
)

type Participant struct {
	gorm.Model
	UserID  *uint
	User    modelsUser.User `gorm:"foreignKey:UserID"`
	EventID uint
	Event   modelsEvent.Event `gorm:"foreignKey:EventID"`
	Attend  bool
	Email   string
}
