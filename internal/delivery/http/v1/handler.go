package v1

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
	"github.com/vectornko/organization-api/internal/service"
	"github.com/vectornko/organization-api/pkg/logger"
)

var log = logger.NewLogger()

type Handler struct {
	services *service.Services
	gc       geo.Geocoder
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
		gc:       openstreetmap.Geocoder(),
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitOrganizationRoutes(v1)
	}
}
