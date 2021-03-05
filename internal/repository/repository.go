package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository/postgres"
)

type Organization interface {
	CreateOrganization(m domain.Organization, userId int) (int, error)
}

type Role interface {
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
	}
}
