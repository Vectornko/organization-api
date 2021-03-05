package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository/postgres"
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
	RoleAccess(userId, orgId int, accessType string) (bool, error)
}

// Параметры ролей
const (
	EditOrganization   = "edit_organization"
	DeleteOrganization = "delete_organization"
)

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
