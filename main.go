package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"syncal/internal/events"
	"syncal/internal/events/handler"
)

func main() {

	var repository []events.Event
	handlerEvent := handler.NewHandlerEvent(repository)

	router := gin.Default()

	// Ruta principal que responde "Hola Mundo"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola Mundo")
	})

	router.POST("/events", handlerEvent.Create)

	router.Run()
}
