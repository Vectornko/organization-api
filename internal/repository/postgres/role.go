package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"strings"
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

	fmt.Println(userId, orgId, accessType)

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

func (r *RolePostgres) CreateRole(m domain.Role) (int, error) {
	var orgId int

	query := fmt.Sprintf("INSERT INTO %s (organization_id, name, edit_organization, delete_organization, create_service, edit_service, delete_service, create_role, edit_role, delete_role, create_employee, edit_employee, delete_employee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id", rolesTable)
	row := r.db.QueryRow(query, &m.OrganizationId, &m.Name, &m.EditOrganization, &m.DeleteOrganization, &m.CreateService, &m.EditService, &m.DeleteService, &m.CreateRole, &m.EditRole, &m.DeleteRole, &m.CreateEmployee, &m.EditEmployee, &m.DeleteEmployee)
	if err := row.Scan(&orgId); err != nil {
		return 0, err
	}

	return orgId, nil
}

func (r *RolePostgres) GetAllRoles(orgId int) ([]domain.Role, error) {
	var roles []domain.Role

	query := fmt.Sprintf("SELECT id, organization_id, name, edit_organization, delete_organization, create_service, edit_service, delete_service, create_role, edit_role, delete_role, create_employee, edit_employee, delete_employee FROM %s WHERE organization_id=$1", rolesTable)
	rows, err := r.db.Query(query, orgId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var role domain.Role
		err := rows.Scan(&role.Id, &role.OrganizationId, &role.Name, &role.EditOrganization, &role.DeleteOrganization, &role.CreateService, &role.EditService, &role.DeleteService, &role.CreateRole, &role.EditRole, &role.DeleteRole, &role.CreateEmployee, &role.EditEmployee, &role.DeleteEmployee)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}
	return roles, err
}

func (r *RolePostgres) GetRoleById(roleId int) (domain.Role, error) {
	var role domain.Role

	query := fmt.Sprintf("SELECT id, organization_id, name, edit_organization, delete_organization, create_service, edit_service, delete_service, create_role, edit_role, delete_role, create_employee, edit_employee, delete_employee FROM %s WHERE id=$1", rolesTable)
	row := r.db.QueryRow(query, roleId)
	err := row.Scan(&role.Id, &role.OrganizationId, &role.Name, &role.EditOrganization, &role.DeleteOrganization, &role.CreateService, &role.EditService, &role.DeleteService, &role.CreateRole, &role.EditRole, &role.DeleteRole, &role.CreateEmployee, &role.EditEmployee, &role.DeleteEmployee)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *RolePostgres) UpdateRole(m domain.UpdateRole) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if m.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *m.Name)
		argId++
	}

	if m.EditOrganization != nil {
		setValues = append(setValues, fmt.Sprintf("edit_organization=$%d", argId))
		args = append(args, *m.EditOrganization)
		argId++
	}

	if m.DeleteOrganization != nil {
		setValues = append(setValues, fmt.Sprintf("delete_organization=$%d", argId))
		args = append(args, *m.DeleteOrganization)
		argId++
	}

	if m.CreateService != nil {
		setValues = append(setValues, fmt.Sprintf("create_service=$%d", argId))
		args = append(args, *m.CreateService)
		argId++
	}

	if m.EditService != nil {
		setValues = append(setValues, fmt.Sprintf("edit_service=$%d", argId))
		args = append(args, *m.EditService)
		argId++
	}

	if m.DeleteService != nil {
		setValues = append(setValues, fmt.Sprintf("delete_service=$%d", argId))
		args = append(args, *m.DeleteService)
		argId++
	}

	if m.CreateRole != nil {
		setValues = append(setValues, fmt.Sprintf("create_role=$%d", argId))
		args = append(args, *m.CreateRole)
		argId++
	}

	if m.EditRole != nil {
		setValues = append(setValues, fmt.Sprintf("edit_role=$%d", argId))
		args = append(args, *m.EditRole)
		argId++
	}

	if m.DeleteRole != nil {
		setValues = append(setValues, fmt.Sprintf("delete_role=$%d", argId))
		args = append(args, *m.DeleteRole)
		argId++
	}

	if m.CreateEmployee != nil {
		setValues = append(setValues, fmt.Sprintf("create_employee=$%d", argId))
		args = append(args, *m.CreateEmployee)
		argId++
	}

	if m.EditEmployee != nil {
		setValues = append(setValues, fmt.Sprintf("edit_employee=$%d", argId))
		args = append(args, *m.EditEmployee)
		argId++
	}

	if m.DeleteEmployee != nil {
		setValues = append(setValues, fmt.Sprintf("delete_employee=$%d", argId))
		args = append(args, *m.DeleteEmployee)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=%d`,
		rolesTable, setQuery, m.Id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *RolePostgres) DeleteRole(roleId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", rolesTable)
	_, err := r.db.Exec(query, roleId)
	return err
}
