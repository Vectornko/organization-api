package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"strings"
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

func (r *OrganizationPostgres) IsEnable(orgId int) (bool, error) {
	var enabled bool

	query := fmt.Sprintf("SELECT is_enable FROM %s WHERE id=$1", organizationsTable)
	row := r.db.QueryRow(query, orgId)
	err := row.Scan(&enabled)
	if err != nil {
		return false, err
	}

	return enabled, nil
}

func (r *OrganizationPostgres) GetAllOrganizations() ([]domain.Organization, error) {
	var organizations []domain.Organization

	query := fmt.Sprintf("SELECT id, name, description, email, phone, site, coordinates, office, date_creation, date_update, is_active, is_enable FROM %s WHERE is_enable=true", organizationsTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var organization domain.Organization
		err := rows.Scan(&organization.Id, &organization.Name, &organization.Description, &organization.Email, &organization.Phone, &organization.Site, &organization.Coordinates, &organization.Office, &organization.DateCreation, &organization.DateUpdate, &organization.IsActive, &organization.IsEnable)
		if err != nil {
			return nil, err
		}

		organizations = append(organizations, organization)
	}
	return organizations, err
}

func (r *OrganizationPostgres) GetOrganizationById(orgId int) (domain.Organization, error) {
	var organization domain.Organization

	query := fmt.Sprintf("SELECT id, name, description, email, phone, site, coordinates, office, date_creation, date_update, is_active, is_enable FROM %s WHERE id=$1", organizationsTable)
	row := r.db.QueryRow(query, orgId)
	err := row.Scan(&organization.Id, &organization.Name, &organization.Description, &organization.Email, &organization.Phone, &organization.Site, &organization.Coordinates, &organization.Office, &organization.DateCreation, &organization.DateUpdate, &organization.IsActive, &organization.IsEnable)
	if err != nil {
		return organization, err
	}

	return organization, nil
}

func (r *OrganizationPostgres) UpdateOrganization(m domain.UpdateOrganization) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if m.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *m.Name)
		argId++
	}

	if m.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *m.Description)
		argId++
	}

	if m.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *m.Email)
		argId++
	}

	if m.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *m.Phone)
		argId++
	}

	if m.Site != nil {
		setValues = append(setValues, fmt.Sprintf("site=$%d", argId))
		args = append(args, *m.Site)
		argId++
	}

	if m.Coordinates != nil {
		setValues = append(setValues, fmt.Sprintf("coordinates=$%d", argId))
		args = append(args, *m.Coordinates)
		argId++
	}

	if m.Office != nil {
		setValues = append(setValues, fmt.Sprintf("office=$%d", argId))
		args = append(args, *m.Office)
		argId++
	}

	if m.IsActive != nil {
		setValues = append(setValues, fmt.Sprintf("is_active=$%d", argId))
		args = append(args, *m.IsActive)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=%d`,
		organizationsTable, setQuery, m.Id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *OrganizationPostgres) DeleteOrganization(orgId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", organizationsTable)
	_, err := r.db.Exec(query, orgId)
	return err
}
