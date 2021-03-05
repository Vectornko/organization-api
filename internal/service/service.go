package service

import (
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository"
)

type Organization interface {
	CreateOrganization(m domain.Organization, userId int) (int, error)
}

type Role interface {
}

type Employee interface {
}

type Services struct {
	Organization
	Role
	Employee
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Organization: NewOrganizationPostgres(repo),
	}
}
