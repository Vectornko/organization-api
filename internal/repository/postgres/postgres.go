package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

// Список таблиц
const (
	organizationsTable         = "organizations"
	rolesTable                 = "roles"
	organizationsUsersTable    = "organizations_users"
	organizationDocumentsTable = "organization_documents"
)

// ID ролей
const (
	orgCreatorRole = "1"
)

// Конфиг для postgres
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Подключени к postgresSQL
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
