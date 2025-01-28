package notificador

import (
	"syncal/database"
	"syncal/internal/notificador/service"
)

var Notificar = func() service.MailService {
	var mailService service.MailService
	database.Database.First(&mailService)
	return mailService
}()
