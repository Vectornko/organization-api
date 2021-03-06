package service

import (
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository"
)

type Organization interface {
	CreateOrganization(m domain.Organization, userId int) (int, error)
	IsEnable(orgId int) (bool, error)
	GetAllOrganizations() ([]domain.Organization, error)
	GetOrganizationById(orgId int) (domain.Organization, error)
	UpdateOrganization(m domain.UpdateOrganization, userId int) error
	DeleteOrganization(orgId, userId int) error
}

type Role interface {
	CreateRole(m domain.Role, userId int) (int, error)
	GetAllRoles(orgId int) ([]domain.Role, error)
	GetRoleById(roleId int) (domain.Role, error)
	UpdateRole(m domain.UpdateRole, userId int) error
	DeleteRole(orgId, roleId, userId int) error

	RoleAccess(userId, orgId int, accessType string) (bool, error)
}

type Employee interface {
	CreateEmployee(m domain.OrganizationsUsers, userId int) (int, error)
	GetAllEmployees(orgId int) ([]domain.OrganizationsUsers, error)
	GetEmployeeById(employeeId int) (domain.OrganizationsUsers, error)
	UpdateEmployee(m domain.UpdateEmployee, userId int) error
	DeleteEmployee(orgId, employeeId, userId int) error
}

type Services struct {
	Organization
	Role
	Employee
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Organization: NewOrganizationPostgres(repo),
		Role:         NewRoleService(repo),
		Employee:     NewEmployeeService(repo),
	}
}
