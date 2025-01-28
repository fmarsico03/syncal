package service

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	SMTPHost     string
	SMTPPort     int
	SenderEmail  string
	SenderPasswd string
}

func (ms *MailService) Send(mail Mailer) error {
	// Construir el correo.
	subject, body := mail.Build()
	recipient := mail.Receiver()

	// Crear el mensaje.
	m := gomail.NewMessage()
	m.SetHeader("From", ms.SenderEmail)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Configuración del servidor SMTP.
	d := gomail.NewDialer(ms.SMTPHost, ms.SMTPPort, ms.SenderEmail, ms.SenderPasswd)

	// Enviar el correo.
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send utilEmail: %w", err)
	}

	return nil
}
