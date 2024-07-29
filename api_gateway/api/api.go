package api

import (
	"api_gateway/api/handler"
	"api_gateway/service"

	"github.com/gin-gonic/gin"
)

func Api() {
	router := gin.Default()

	services := service.Service()

	h := handler.NewHandler(services)

	api := router.Group("/api")

	api.POST("/crt", h.CreateContent)
	api.GET("/contents", h.GetContents)
	api.GET("/content/:id", h.GetContentById)
	api.PUT("/update/:id",h.UpdateContent)
	api.DELETE("/delete/:id",h.Delete)


	router.Run(":8080")

}
