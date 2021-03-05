package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
)

type OrganizationPostgres struct {
	db *sqlx.DB
}

func NewOrganizationPostgres(db *sqlx.DB) *OrganizationPostgres {
	return &OrganizationPostgres{db: db}
}

func (r *OrganizationPostgres) CreateOrganization(m domain.Organization, userId int) (int, error) {
	var orgId int

	tx, _ := r.db.Begin()

	query := fmt.Sprintf("INSERT INTO %s (name, email, phone, coordinates) VALUES ($1, $2, $3, $4) RETURNING id", organizationsTable)
	row := tx.QueryRow(query, &m.Name, &m.Email, &m.Phone, &m.Coordinates)
	if err := row.Scan(&orgId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (organization_id, user_id, role_id) VALUES ($1, $2, $3)", organizationsUsersTable)
	_, err := tx.Exec(query, orgId, userId, orgCreatorRole)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return orgId, nil
}
