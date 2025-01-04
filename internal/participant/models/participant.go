package models

import (
	modelsEvent "syncal/internal/events/models"
	modelsUser "syncal/internal/users/models"
)

type Participant struct {
	UserID  uint              `gorm:"primaryKey"` // Referencia a la clave primaria de User
	User    modelsUser.User   `gorm:"foreignKey:UserID"`
	EventID uint              `gorm:"primaryKey"` // Referencia a la clave primaria de Event
	Event   modelsEvent.Event `gorm:"foreignKey:EventID"`
	Attend  bool
}
