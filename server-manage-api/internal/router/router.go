package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server-manage-api/internal/handlers"
	"server-manage-api/internal/services"
)

func Setup(router *gin.Engine, services *services.Services) {

	router.Use(cors.Default())

	h := handlers.NewHandler(services.ServerManagerService)

	router.GET("/health", handlers.Health)
	serversGroup := router.Group("/servers")
	{
		serversGroup.GET("/", h.ListServers)
		serversGroup.POST("/", h.CreateServer)
		serversGroup.GET("/:id", h.GetServerByID)
		serversGroup.DELETE("/:id", h.DeleteServerByID)
	}

	router.NoRoute(handlers.NoRoute)
}
