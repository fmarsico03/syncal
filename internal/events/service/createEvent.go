package service

import (
	"fmt"
	"syncal/internal/events"
	"syncal/internal/events/request"
	"syncal/internal/events/response"
	"syncal/internal/users"
)

func CreateEvent(req request.CreateEventRequest) (response.CreateEventResponse, events.Event) {
	franco := users.NewUser("Franco", "Marsico", "fmarsico03@gmail.com")

	event := events.NewEventComplete(req.Title, *franco, req.Start, req.End)
	if &req.Description != nil {
		event.SetDescription(req.Description)
	}
	if &req.Location != nil {
		event.SetLocation(req.Location)
	}
	if &req.MeetLink != nil {
		event.SetMeetLink(req.MeetLink)
	}

	fmt.Printf("Titulo: %s, Descripcion: %s, Meet: %s\n", event.Title(), event.Description(), event.MeetLink())

	event.SetCreatedBy(*franco)
	var participants []string
	participants = req.Participants
	//Logica
	//todo Buscar en BD los participantes -> Si no existe alguno enviar ademas un msj de error
	// O habria que ver si existe el mail? Tantear eso. Luego mapear varios objetos participantes.

	return *response.NewCreateEventResponse(
			event.CreatedBy().Mail(),
			event.Description(),
			event.End(),
			event.Location(),
			event.MeetLink(),
			participants,
			event.Start(),
			event.Title()),
		*event
}
