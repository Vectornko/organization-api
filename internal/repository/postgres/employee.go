package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vectornko/organization-api/internal/domain"
	"strings"
)

type EmployeePostgres struct {
	db *sqlx.DB
}

func NewEmployeePostgres(db *sqlx.DB) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}

func (r *EmployeePostgres) CreateEmployee(m domain.OrganizationsUsers) (int, error) {
	var employeeId int

	query := fmt.Sprintf("INSERT INTO %s (organization_id, user_id, role_id) VALUES ($1, $2, $3) RETURNING id", organizationsUsersTable)
	row := r.db.QueryRow(query, &m.OrganizationId, &m.UserId, &m.RoleId)
	if err := row.Scan(&employeeId); err != nil {
		if err.Error() == "pq: insert or update on table \"organizations_users\" violates foreign key constraint \"organizations_users_role_id_fkey\"" {
			return 0, errors.New("no role with this id")
		}
		return 0, err
	}

	return employeeId, nil
}

func (r *EmployeePostgres) EmployeeExist(employeeId, orgId int) (bool, error) {
	var roleId int
	query := fmt.Sprintf("SELECT role_id FROM %s WHERE user_id=$1 AND organization_id=$2", organizationsUsersTable)
	row := r.db.QueryRow(query, employeeId, orgId)
	err := row.Scan(&roleId)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *EmployeePostgres) EmployeeConfirmed(employeeId, orgId int) (bool, error) {
	var confirmed bool
	query := fmt.Sprintf("SELECT confirmed FROM %s WHERE user_id=$1 AND organization_id=$2", organizationsUsersTable)
	row := r.db.QueryRow(query, employeeId, orgId)
	err := row.Scan(&confirmed)
	if err != nil {
		return false, err
	}

	return confirmed, nil
}

func (r *EmployeePostgres) GetAllEmployees(orgId int) ([]domain.OrganizationsUsers, error) {
	var employees []domain.OrganizationsUsers

	query := fmt.Sprintf("SELECT id, organization_id, user_id, role_id, confirmed FROM %s WHERE organization_id=$1", organizationsUsersTable)
	rows, err := r.db.Query(query, orgId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee domain.OrganizationsUsers
		err := rows.Scan(&employee.Id, &employee.OrganizationId, &employee.UserId, &employee.RoleId, &employee.Confirmed)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}
	return employees, err
}

func (r *EmployeePostgres) GetEmployeeById(employeeId int) (domain.OrganizationsUsers, error) {
	var employee domain.OrganizationsUsers

	query := fmt.Sprintf("SELECT id, organization_id, user_id, role_id, confirmed FROM %s WHERE id=$1", organizationsUsersTable)
	row := r.db.QueryRow(query, employeeId)
	err := row.Scan(&employee.Id, &employee.OrganizationId, &employee.UserId, &employee.RoleId, &employee.Confirmed)
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (r *EmployeePostgres) UpdateEmployee(m domain.UpdateEmployee) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if m.RoleId != nil {
		setValues = append(setValues, fmt.Sprintf("role_id=$%d", argId))
		args = append(args, *m.RoleId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=%d`,
		organizationsUsersTable, setQuery, m.Id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *EmployeePostgres) DeleteEmployee(employeeId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", organizationsUsersTable)
	_, err := r.db.Exec(query, employeeId)
	return err
}
