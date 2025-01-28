package service

import (
	"fmt"
	"syncal/database"
	modelsEvent "syncal/internal/events/models"
	"syncal/internal/notificador"
	"syncal/internal/notificador/service"
	"syncal/internal/participant/models"
	"syncal/internal/utils/utilEmail"
	"time"
)

func CreateParticipants(event modelsEvent.Event, mail string, endGroupTime time.Time) {
	var participant models.Participant
	user, err := utilEmail.FindUserByEmail(mail)
	if err != nil {
		participant.UserID = nil
	} else {
		participant.UserID = &user.ID
	}
	participant.Event = event
	participant.Email = mail
	database.Database.Create(&participant)

	invitationMailer := mapEventToInvitation(event, mail, endGroupTime)
	err = notificador.Notificar.Send(&invitationMailer)
	if err != nil {
		fmt.Println(err)
	}
}

func mapEventToInvitation(event modelsEvent.Event, to string, endGroupTime time.Time) service.InvitationMailer {
	return service.InvitationMailer{
		Title:        event.Title,
		Name:         event.CreatedBy.Name,
		LastName:     event.CreatedBy.Lastname,
		EmailCreator: event.CreatedBy.Mail,
		Location:     event.Location,
		Meet:         event.MeetLink,
		Start:        event.Start,
		End:          event.End,
		EndGroup:     endGroupTime,
		Type:         event.Type,
		Description:  event.Description,
		To:           to,
	}
}
