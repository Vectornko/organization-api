package service

import (
	"errors"
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

func (s *OrganizationService) IsEnable(orgId int) (bool, error) {
	return s.repo.Organization.IsEnable(orgId)
}

func (s *OrganizationService) GetAllOrganizations() ([]domain.Organization, error) {
	return s.repo.Organization.GetAllOrganizations()
}

func (s *OrganizationService) GetOrganizationById(orgId int) (domain.Organization, error) {
	return s.repo.Organization.GetOrganizationById(orgId)
}

func (s *OrganizationService) UpdateOrganization(m domain.UpdateOrganization, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, m.Id, repository.EditOrganization)
	if err != nil {
		return err
	}

	if access == true {
		return s.repo.Organization.UpdateOrganization(m)
	} else {
		return errors.New("access denied")
	}
}

func (s *OrganizationService) DeleteOrganization(orgId, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, orgId, repository.DeleteOrganization)
	if err != nil {
		return err
	}

	if access == true {
		return s.repo.Organization.DeleteOrganization(orgId)
	} else {
		return errors.New("access denied")
	}
}
