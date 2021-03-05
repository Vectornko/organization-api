package repository

import (
	"github.com/jmoiron/sqlx"
)

type Organization interface {
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
	return &Repository{}
}
