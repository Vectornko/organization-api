package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RolePostgres struct {
	db *sqlx.DB
}

func NewRolePostgres(db *sqlx.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) RoleAccess(userId, orgId int, accessType string) (bool, error) {
	var access bool
	var roleId int

	query := fmt.Sprintf("SELECT role_id FROM %s WHERE user_id=$1 AND organization_id=$2", organizationsUsersTable)
	row := r.db.QueryRow(query, userId, orgId)
	err := row.Scan(&roleId)
	if err != nil {
		return false, err
	}

	query = fmt.Sprintf("SELECT %s FROM %s WHERE id=$1", accessType, rolesTable)
	row = r.db.QueryRow(query, roleId)
	err = row.Scan(&access)
	if err != nil {
		return false, err
	}

	return access, nil
}
