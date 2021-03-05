package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vectornko/organization-api/internal/domain"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) InitOrganizationRoutes(v1 *gin.RouterGroup) {
	organization := v1.Group("/organization", h.userIdentify)
	{
		organization.POST("/", h.createOrganization)
	}
}

// @Summary Регистрация организации
// @Security ApiKeyAuth
// @Tags auth
// @Description Сервер создаёт организацию и возврощает её id
// @ID createOrganization
// @Accept json
// @Produce json
// @Param input body registrationForm true "Описание"
// @Success 200 {string} UserId
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/ [post]
func (h *Handler) createOrganization(c *gin.Context) {
	// Достаём ID пользователя
	userId, _ := strconv.Atoi(c.GetString(userCtx))
	// Переменная для Body от запроса
	var input registrationForm
	var org domain.Organization

	// Десериализируем Body в нашу структуру
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// Валидация email
	if len(strings.Split(input.Email, "@")) != 2 {
		newErrorResponse(c, http.StatusBadRequest, "invalid email")
		return
	}

	// Валидация номера
	if strings.Split(input.Phone, "")[0] != "+" || strings.Split(input.Phone, "")[1] != "7" {
		newErrorResponse(c, http.StatusBadRequest, "number must start with +7")
		return
	}
	if len(strings.Split(input.Phone, "")) != 12 {
		newErrorResponse(c, http.StatusBadRequest, "invalid number")
		return
	}

	// Валидация координат
	_, err = h.gc.ReverseGeocode(input.Address.Latitude, input.Address.Longitude)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid coordinates")
		return
	}

	// Преобразовываем input в нужную нам структуру
	arr := []string{fmt.Sprintf("%f", input.Address.Latitude), fmt.Sprintf("%f", input.Address.Longitude)}
	org.Coordinates = strings.Join(arr, ",")
	org.Name = input.Name
	org.Email = input.Email
	org.Phone = input.Phone

	// Вызываем метод слоя Service
	orgId, err := h.services.Organization.CreateOrganization(org, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Отдаём ответ в виде JSON
	c.JSON(http.StatusOK, orgId)
}

type registrationForm struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Address address `json:"address" binding:"required"`
}

type address struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
