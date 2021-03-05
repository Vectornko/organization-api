package service

import (
	"errors"
	"github.com/vectornko/organization-api/internal/domain"
	"github.com/vectornko/organization-api/internal/repository"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type EmployeeService struct {
	repo *repository.Repository
}

func NewEmployeeService(repo *repository.Repository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(m domain.OrganizationsUsers, userId int) (int, error) {
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!СДЕЛАТЬ ОТПАРВКУ письма на email!!!!!!!!!!!!!!!!!
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	access, err := s.repo.Role.RoleAccess(userId, m.OrganizationId, repository.CreateEmployee)
	if err != nil {
		return 0, errors.New("access denied")
	}

	arr := []string{"http://localhost:8000/api/v1/", strconv.Itoa(m.UserId)}
	req := strings.Join(arr, "")
	resp, err := http.Get(req)
	if err != nil {
		return 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if string(body) == "false" {
		return 0, errors.New("user does not exist")
	}

	exist, err := s.repo.Employee.EmployeeExist(m.UserId, m.OrganizationId)
	if err != nil {
		return 0, err
	}

	if exist == true {
		return 0, errors.New("the user is already an employee")
	}

	if access == true {
		return s.repo.Employee.CreateEmployee(m)
	} else {
		return 0, errors.New("access denied")
	}
	return 0, nil
}

func (s *EmployeeService) GetAllEmployees(orgId int) ([]domain.OrganizationsUsers, error) {
	return s.repo.Employee.GetAllEmployees(orgId)
}

func (s *EmployeeService) GetEmployeeById(employeeId int) (domain.OrganizationsUsers, error) {
	return s.repo.Employee.GetEmployeeById(employeeId)
}

func (s *EmployeeService) UpdateEmployee(m domain.UpdateEmployee, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, m.OrganizationId, repository.EditRole)
	if err != nil {
		return errors.New("access denied")
	}

	if access == true {
		return s.repo.Employee.UpdateEmployee(m)
	} else {
		return errors.New("access denied")
	}
}

func (s *EmployeeService) DeleteEmployee(orgId, employeeId, userId int) error {
	access, err := s.repo.Role.RoleAccess(userId, orgId, repository.DeleteRole)
	if err != nil {
		return errors.New("access denied")
	}

	if access == true {
		return s.repo.Employee.DeleteEmployee(employeeId)
	} else {
		return errors.New("access denied")
	}
}
