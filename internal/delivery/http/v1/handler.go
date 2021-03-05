package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vectornko/organization-api/internal/service"
	"github.com/vectornko/organization-api/pkg/logger"
)

var log = logger.NewLogger()

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitOrganizationRoutes(v1)
	}
}
