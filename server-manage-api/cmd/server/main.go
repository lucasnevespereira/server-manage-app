package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"server-manage-api/configs"
	_ "server-manage-api/docs"
	"server-manage-api/internal/router"
	"server-manage-api/internal/services"
)

func main() {
	r := gin.Default()
	config := configs.Load()
	services := services.InitServices(config)
	router.Setup(r, services)
	url := ginSwagger.URL(fmt.Sprintf("%s:%d/swagger/doc.json", config.SwaggerHost, config.Port))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
