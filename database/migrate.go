package database

import (
	modelsEvent "syncal/internal/events/models"
	modelsParticipant "syncal/internal/participant/models"
	modelsUser "syncal/internal/users/models"
)

func Migrate() {
	Database.AutoMigrate(modelsEvent.Event{})
	Database.AutoMigrate(modelsUser.User{})
	Database.AutoMigrate(modelsParticipant.Participant{})
}
