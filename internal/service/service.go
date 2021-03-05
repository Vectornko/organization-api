package service

import (
	"github.com/vectornko/organization-api/internal/repository"
)

type Organization interface {
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
	return &Services{}
}
