package events

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	subRouter := router.Group("/events")

	subRouter.POST("/", Create)
}
