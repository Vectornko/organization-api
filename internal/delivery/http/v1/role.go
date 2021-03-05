package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vectornko/organization-api/internal/domain"
	"net/http"
	"strconv"
)

func (h *Handler) InitRoleRoutes(role *gin.RouterGroup) {
	role.POST("/", h.createRole)
	role.GET("/", h.getAllRoles)
	role.GET("/:role_id", h.getRoleById)
	role.PUT("/:role_id", h.updatedRole)
	role.DELETE("/:role_id", h.deleteRole)
}

// @Summary Создание роли
// @Security ApiKeyAuth
// @Tags role
// @Description Сервер создаёт роль для определённой организации
// @ID createRole
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param input body domain.Role true "id и organization_id вводить не нужно"
// @Success 200 {int} RoleId
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/role/ [post]
func (h *Handler) createRole(c *gin.Context) {
	var input domain.Role
	userId, _ := strconv.Atoi(c.GetString(userCtx))
	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusOK, "invalid organization id")
		return
	}

	err = c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusOK, "invalid input body")
		return
	}

	input.OrganizationId = orgId

	id, err := h.services.Role.CreateRole(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary Список всех ролей организации
// @Security ApiKeyAuth
// @Tags role
// @Description Сервер отдаёт список ролей организации
// @ID getAllRoles
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Success 200 {string} Roles
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/role/ [get]
func (h *Handler) getAllRoles(c *gin.Context) {
	var roles Roles
	var err error

	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusOK, "invalid organization id")
		return
	}

	roles.Roles, err = h.services.Role.GetAllRoles(orgId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, roles)
}

// @Summary Информация о роли
// @Security ApiKeyAuth
// @Tags role
// @Description Сервер возвращает роль организации по id
// @ID getRoleById
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  role_id path int true "id роли"
// @Success 200 {object} domain.Role
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/role/{role_id} [get]
func (h *Handler) getRoleById(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	role, err := h.services.Role.GetRoleById(roleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Summary Изменить роль
// @Security ApiKeyAuth
// @Tags role
// @Description Сервер изменяет данные роли
// @ID updateOrganization
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  role_id path int true "id роли"
// @Param input body domain.UpdateRole true "id и organization_id вводить не нужно"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/role/{role_id} [put]
func (h *Handler) updatedRole(c *gin.Context) {
	var input domain.UpdateRole
	userId, _ := strconv.Atoi(c.GetString(userCtx))

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	input.OrganizationId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	input.Id, err = strconv.Atoi(c.Param("role_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	err = h.services.Role.UpdateRole(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "role update")
}

// @Summary Удалить роль
// @Security ApiKeyAuth
// @Tags role
// @Description Сервер удаляет роль
// @ID deleteRole
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  role_id path int true "id роли"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/role/{role_id} [delete]
func (h *Handler) deleteRole(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetString(userCtx))

	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	roleId, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id")
		return
	}

	err = h.services.Role.DeleteRole(orgId, roleId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "role deleted")
}

type Roles struct {
	Roles []domain.Role `json:"roles"`
}
