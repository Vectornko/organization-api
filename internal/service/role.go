package service

import (
	"errors"
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository"
)

type RoleService struct {
	repo *repository.Repository
}

func NewRoleService(repo *repository.Repository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) CreateRole(m domain.Role, userId int) (int, error) {
	access, err := s.repo.Role.RoleAccess(userId, m.OrganizationId, repository.CreateRole)
	if err != nil {
		return 0, errors.New("access denied")
	}

	if access == true {
		return s.repo.Role.CreateRole(m)
	} else {
		return 0, errors.New("access denied")
	}
}

func (s *RoleService) GetAllRoles(orgId int) ([]domain.Role, error) {
	return s.repo.Role.GetAllRoles(orgId)
}

func (s *RoleService) GetRoleById(roleId int) (domain.Role, error) {
	role, err := s.repo.Role.GetRoleById(roleId)
	if err != nil {
		return role, errors.New("role missing")
	}
	return role, err
}

func (s *RoleService) UpdateRole(m domain.UpdateRole, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, m.OrganizationId, repository.EditRole)
	if err != nil {
		return errors.New("access denied")
	}

	if access == true {
		return s.repo.Role.UpdateRole(m)
	} else {
		return errors.New("access denied")
	}
}

func (s *RoleService) DeleteRole(orgId, roleId, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, orgId, repository.DeleteRole)
	if err != nil {
		return errors.New("access denied")
	}

	if access == true {
		return s.repo.Role.DeleteRole(roleId)
	} else {
		return errors.New("access denied")
	}
}

func (s *RoleService) RoleAccess(userId, orgId int, accessType string) (bool, error) {
	return s.repo.Role.RoleAccess(userId, orgId, accessType)
}
