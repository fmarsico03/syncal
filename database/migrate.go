package database

import (
	modelsEvent "syncal/internal/events/models"
	"syncal/internal/notificador/service"
	modelsParticipant "syncal/internal/participant/models"
	modelsUser "syncal/internal/users/models"
)

func Migrate() {
	Database.AutoMigrate(modelsEvent.Event{})
	Database.AutoMigrate(modelsUser.User{})
	Database.AutoMigrate(modelsParticipant.Participant{})
	Database.AutoMigrate(service.MailService{})
}
