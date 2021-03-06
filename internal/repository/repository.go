package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository/postgres"
)

// Параметры ролей
const (
	EditOrganization   = "edit_organization"
	DeleteOrganization = "delete_organization"
	CreateRole         = "create_role"
	EditRole           = "edit_role"
	DeleteRole         = "delete_role"
	CreateEmployee     = "create_employee"
	EditEmployee       = "edit_employee"
	DeleteEmployee     = "delete_employee"
	CreateService      = "create_service"
	EditService        = "edit_service"
	DeleteService      = "delete_service"
)

type Organization interface {
	CreateOrganization(m domain.Organization, userId int) (int, error)
	GetAllOrganizations() ([]domain.Organization, error)
	GetOrganizationById(orgId int) (domain.Organization, error)
	UpdateOrganization(m domain.UpdateOrganization) error
	DeleteOrganization(orgId int) error

	IsEnable(orgId int) (bool, error)
}

type Role interface {
	CreateRole(m domain.Role) (int, error)
	GetAllRoles(orgId int) ([]domain.Role, error)
	GetRoleById(roleId int) (domain.Role, error)
	UpdateRole(m domain.UpdateRole) error
	DeleteRole(roleId int) error

	RoleAccess(userId, orgId int, accessType string) (bool, error)
}

type Employee interface {
	CreateEmployee(m domain.OrganizationsUsers) (int, error)
	GetAllEmployees(orgId int) ([]domain.OrganizationsUsers, error)
	GetEmployeeById(roleId int) (domain.OrganizationsUsers, error)
	UpdateEmployee(m domain.UpdateEmployee) error
	DeleteEmployee(employeeId int) error

	EmployeeExist(employeeId, orgId int) (bool, error)
	EmployeeConfirmed(employeeId, orgId int) (bool, error)
}

type Repository struct {
	Organization
	Role
	Employee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Organization: postgres.NewOrganizationPostgres(db),
		Role:         postgres.NewRolePostgres(db),
		Employee:     postgres.NewEmployeePostgres(db),
	}
}
