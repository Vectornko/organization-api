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
