package main

import (
	"github.com/gin-gonic/gin"
	"syncal/internal/events"
)

func main() {

	router := gin.Default()

	//events
	events.RegisterRoutes(router)
	//migration

	//users
	//Routes and migration

	err := router.Run()
	if err != nil {
		return
	}

}
