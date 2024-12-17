package events

type HandlerEvent struct{}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{}
}

/*
func (h *HandlerEvent) Create(c main.CreateEventRequest) {
	event := NewEventComplete(req.Title, req.CreatedBy, req.Frequency...)

	// Agregar los campos opcionales
	if req.Location != "" {
		event.SetLocation(req.Location)
	}
	if req.Description != "" {
		event.SetDescription(req.Description)
	}
	if req.MeetLink != "" {
		event.SetMeetLink(req.MeetLink)
	}
}
*/
