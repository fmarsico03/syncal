package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/internal/events"
	"syncal/internal/users"
	"time"
)

func main() {

	franco := users.NewUser("Franco", "Marsico", "fmarsico03@gmail.com")
	var repositotyEvents []events.Event
	// Crear una nueva instancia de Gin
	router := gin.Default()

	// Ruta principal que responde "Hola Mundo"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola Mundo")
	})

	router.POST("/events", func(c *gin.Context) {
		var frequencys []events.Recurrence

		// Parsear y validar el cuerpo
		var req events.CreateEventRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, f := range req.Frequency {
			frequency := events.NewRecurrence(f.GetStartedAt(), f.GetFinishedAt(), events.DayOfWeek(f.GetDayOfWeek()))
			frequencys = append(frequencys, *frequency)
			fmt.Printf("Frecuencia: Day: %d, Start: %s, End: %s\n", f.DayOfWeek, f.StartedAt.Format(time.RFC3339), f.FinishedAt.Format(time.RFC3339))
		}

		// Crear el evento utilizando el constructor obligatorio
		event := events.NewEventComplete(req.Title, *franco, frequencys...)

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

		repositotyEvents = append(repositotyEvents, *event)

		// Retornar una respuesta de éxito
		c.JSON(http.StatusOK, gin.H{
			"message": "Event created successfully",
			"event": gin.H{
				"title":       event.Title(),
				"created_by":  event.CreatedBy().Name(),
				"location":    event.Location(),
				"description": event.Description(),
				"meet_link":   event.MeetLink(),
				"frequency":   event.Frequency(),
			},
		})
	})

	router.GET("/events", func(c *gin.Context) {
		var eventsList []gin.H

		// Recorremos los eventos y formamos el slice de eventos en formato JSON
		for _, event := range repositotyEvents {
			var eventFrequencies []gin.H
			for _, freq := range event.Frequency() {
				// Convirtiendo cada frecuencia a formato adecuado
				eventFrequencies = append(eventFrequencies, freq.ToJSON())
			}
			eventData := gin.H{
				"title":       event.Title(),
				"created_by":  event.CreatedBy().Name(),
				"location":    event.Location(),
				"description": event.Description(),
				"meet_link":   event.MeetLink(),
				"frequency":   eventFrequencies,
			}

			// Agregamos cada evento al slice
			eventsList = append(eventsList, eventData)
		}

		// Respondemos con los eventos en formato JSON
		c.JSON(http.StatusOK, gin.H{"events": eventsList})
	})

	router.Run()
}
