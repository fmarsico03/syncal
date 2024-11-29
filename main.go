package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear una nueva instancia de Gin
	router := gin.Default()

	// Ruta principal que responde "Hola Mundo"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola Mundo")
	})

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hola Mundo")
	})

	router.Run()
}
