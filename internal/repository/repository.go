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
	}
}
