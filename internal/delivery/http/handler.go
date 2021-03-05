package http

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/vectornko/organization-api/docs"
	v1 "github.com/vectornko/organization-api/internal/delivery/http/v1"
	"github.com/vectornko/organization-api/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitAPI() *gin.Engine {
	router := gin.New()

	handlerV1 := v1.NewHandler(h.services)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	return router
}
