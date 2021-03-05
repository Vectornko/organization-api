package service

import (
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository"
)

type OrganizationService struct {
	repo *repository.Repository
}

func NewOrganizationPostgres(repo *repository.Repository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(m domain.Organization, userId int) (int, error) {
	return s.repo.Organization.CreateOrganization(m, userId)
}
