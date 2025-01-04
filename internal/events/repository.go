package events

import (
	db "syncal/database"
	"syncal/internal/events/models"
)

func MigrarEvent() {
	db.Database.AutoMigrate(models.Event{})
}
