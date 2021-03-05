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
		organization.GET("/", h.getAllOrganizations)
		organization.GET("/:id", h.getOrganizationById)
		organization.PUT("/:id", h.updateOrganization)
		organization.DELETE("/:id", h.deleteOrganization)

		roles := organization.Group(":id/role", h.organizationEnable)
		{
			h.InitRoleRoutes(roles)
		}

		employee := organization.Group(":id/employee", h.organizationEnable)
		{
			h.InitEmployeeRoutes(employee)
		}
	}
}

// @Summary Регистрация организации
// @Security ApiKeyAuth
// @Tags organization
// @Description Сервер создаёт организацию и возврощает её id
// @ID createOrganization
// @Accept json
// @Produce json
// @Param input body registrationForm true "Описание"
// @Success 200 {int} OrganizationId
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

// @Summary Список всех организаций
// @Security ApiKeyAuth
// @Tags organization
// @Description Сервер возвращает список всех организаций
// @ID getAllOrganizations
// @Accept json
// @Produce json
// @Success 200 {object} Organizations
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/ [get]
func (h *Handler) getAllOrganizations(c *gin.Context) {
	var organizations Organizations
	var err error

	organizations.Organizations, err = h.services.Organization.GetAllOrganizations()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organizations)
}

// @Summary Информация о организации
// @Security ApiKeyAuth
// @Tags organization
// @Description Сервер возвращает организацию по id
// @ID getOrganizationById
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Success 200 {object} domain.Organization
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id} [get]
func (h *Handler) getOrganizationById(c *gin.Context) {
	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	organization, err := h.services.Organization.GetOrganizationById(orgId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, organization)
}

// @Summary Изменить организацию
// @Security ApiKeyAuth
// @Tags organization
// @Description Сервер изменяет данные организации
// @ID updateOrganization
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param input body domain.UpdateOrganization true "Описание"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id} [put]
func (h *Handler) updateOrganization(c *gin.Context) {
	var input domain.UpdateOrganization
	userId, _ := strconv.Atoi(c.GetString(userCtx))

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	err = h.services.Organization.UpdateOrganization(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "organization update")
}

// @Summary Удаление организации
// @Security ApiKeyAuth
// @Tags organization
// @Description Сервер удаляет организацию
// @ID deleteOrganization
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id} [delete]
func (h *Handler) deleteOrganization(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetString(userCtx))
	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	err = h.services.Organization.DeleteOrganization(orgId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "organization deleted")
}

type Organizations struct {
	Organizations []domain.Organization `json:"organizations"`
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
