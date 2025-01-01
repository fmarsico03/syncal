package service

import (
	"fmt"
	"syncal/internal/events/models"
	"syncal/internal/events/request"
	"syncal/internal/events/response"
	models2 "syncal/internal/users/models"
)

func CreateEvent(req request.CreateEventRequest) (response.CreateEventResponse, models.Event) {
	franco := models2.NewUser("Franco", "Marsico", "fmarsico03@gmail.com")

	event := models.Event{
		Title:       req.Title,
		CreatedBy:   *franco,
		Start:       req.Start,
		End:         req.End,
		Description: req.Description,
		Location:    req.Location,
		MeetLink:    req.MeetLink,
	}

	fmt.Printf("Titulo: %s, Descripcion: %s, Meet: %s\n", event.Title, event.Description, event.MeetLink)

	var participants []string
	participants = req.Participants
	//Logica
	//todo Buscar en BD los participantes -> Si no existe alguno enviar ademas un msj de error
	// O habria que ver si existe el mail? Tantear eso. Luego mapear varios objetos participantes.

	return *response.NewCreateEventResponse(
			event.CreatedBy.Mail(),
			event.Description,
			event.End,
			event.Location,
			event.MeetLink,
			participants,
			event.Start,
			event.Title),
		event
}
