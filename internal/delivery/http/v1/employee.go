package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vectornko/organization-api/internal/domain"
	"net/http"
	"strconv"
)

func (h *Handler) InitEmployeeRoutes(employee *gin.RouterGroup) {
	employee.POST("/", h.createEmployee)
	employee.GET("/", h.getAllEmployees)
	employee.GET("/:employee_id", h.getEmployeeById)
	employee.PUT("/:employee_id", h.updatedEmployee)
	employee.DELETE("/:employee_id", h.deleteEmployee)
}

// @Summary Добавить сотрудника
// @Security ApiKeyAuth
// @Tags employee
// @Description Сервер создаёт сотрудника и отпровляет запрос пользователю на подтверждение
// @ID createEmployee
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param input body registrationEmployee true "id, organization_id и confirmed вводить не нужно"
// @Success 200 {int} EmployeeId
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/employee/ [post]
func (h *Handler) createEmployee(c *gin.Context) {
	var input registrationEmployee
	var employee domain.OrganizationsUsers

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

	employee.OrganizationId = orgId
	employee.UserId = input.UserId
	employee.RoleId = input.RoleId

	id, err := h.services.Employee.CreateEmployee(employee, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary Список всех сотрудников организации
// @Security ApiKeyAuth
// @Tags employee
// @Description Сервер отдаёт список подтверждённых сотрудников
// @ID getAllEmployees
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Success 200 {string} Employees
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/employee/ [get]
func (h *Handler) getAllEmployees(c *gin.Context) {
	var employees Employees
	var err error

	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusOK, "invalid organization id")
		return
	}

	employees.Employees, err = h.services.Employee.GetAllEmployees(orgId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, employees)
}

// @Summary Информация о сотруднике
// @Security ApiKeyAuth
// @Tags employee
// @Description Сервер возвращает сотрудника организации по id
// @ID getEmployeeById
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  employee_id path int true "id сотрудника"
// @Success 200 {object} domain.OrganizationsUsers
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/employee/{employee_id} [get]
func (h *Handler) getEmployeeById(c *gin.Context) {
	employeeId, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusOK, "invalid employee id")
		return
	}

	employee, err := h.services.Employee.GetEmployeeById(employeeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, employee)
}

// @Summary Изменить сотрудника
// @Security ApiKeyAuth
// @Tags employee
// @Description Сервер изменяет данные сотрудника
// @ID updatedEmployee
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  employee_id path int true "id сотрудника"
// @Param input body domain.UpdateEmployee true "Доступно только изминение роли"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/employee/{employee_id} [put]
func (h *Handler) updatedEmployee(c *gin.Context) {
	var input domain.UpdateEmployee
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

	input.Id, err = strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid employee id")
		return
	}

	err = h.services.Employee.UpdateEmployee(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "role update")
}

// @Summary Удалить сотрудника
// @Security ApiKeyAuth
// @Tags employee
// @Description Сервер удаляет сотрудника
// @ID deleteEmployee
// @Accept json
// @Produce json
// @Param  id path int true "id организации"
// @Param  employee_id path int true "id сотрудника"
// @Success 200 {object} ErrorResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/organization/{id}/employee/{employee_id} [delete]
func (h *Handler) deleteEmployee(c *gin.Context) {
	userId, _ := strconv.Atoi(c.GetString(userCtx))

	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid organization id")
		return
	}

	employeeId, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid employee id")
		return
	}

	err = h.services.Employee.DeleteEmployee(orgId, employeeId, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newGoodResponse(c, http.StatusOK, "employee deleted")
}

type registrationEmployee struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

type Employees struct {
	Employees []domain.OrganizationsUsers `json:"employees"`
}
