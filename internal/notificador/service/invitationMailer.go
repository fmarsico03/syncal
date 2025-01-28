package service

import (
	"fmt"
	"syncal/internal/utils/utilRepeatType"
	"time"
)

type InvitationMailer struct {
	Title        string
	Name         string
	LastName     string
	EmailCreator string
	Location     string
	Meet         string
	Start        time.Time
	End          time.Time
	EndGroup     time.Time
	Type         utilRepeatType.RepeatType
	Description  string
	To           string
}

func (m *InvitationMailer) Receiver() string {
	return m.To
}

func (m *InvitationMailer) Build() (string, string) {
	location := m.Location
	if location == "" {
		location = fmt.Sprintf("reunión por Meet (%s)", m.Meet)
	} else {
		location = fmt.Sprintf("reunión en %s", location)
	}

	data := dataBody{
		title:       m.Title,
		description: m.Description,
		location:    location,
		start:       m.Start,
		end:         m.End,
		endGroup:    m.EndGroup,
		typeOf:      m.Type,
		to:          m.To,
		name:        m.Name,
		last:        m.LastName,
	}
	// Construir el cuerpo en HTML.
	htmlBody := data.buildMessage()

	// Asunto del correo.
	subject := fmt.Sprintf("Invitación: %s", m.Title)

	return subject, htmlBody
}
