package main

import (
	"flag"
	"fmt"
	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	configs "github.com/vectornko/organization-api/internal/config"
	"github.com/vectornko/organization-api/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger.NewLogger()

	// Config initialization
	err := configs.InitConfig()
	if err != nil {
		log.Panic("error initializing configs",
			zap.String("package", "main"),
			zap.String("function", "initConfig"),
			zap.Error(err))
	}

	// Parse flags
	SQLMode := flag.String("mode", "", "mode for migrate. up - create table. down - drop table.")
	flag.Parse()
	log.Info("Migration mode", zap.String("mode", *SQLMode))

	// Database connection
	con := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s",
		viper.GetString("db.username"),
		viper.GetString("db.dbname"),
		viper.GetString("db.password"),
		viper.GetString("db.sslmode"),
		viper.GetString("db.host"))

	db, err := sqlx.Connect("postgres", con)
	if err != nil {
		log.Fatal("Error connection to database",
			zap.String("package", "main"),
			zap.String("function", "sqlx.Connect"),
			zap.Error(err))
	}

	// Upload SQL
	organization, err := dotsql.LoadFromFile("internal/schema/organization.sql")
	data, err := dotsql.LoadFromFile("internal/schema/data.sql")
	if err != nil {
		log.Fatal("sql upload error",
			zap.String("package", "main"),
			zap.String("function", "dotsql.LoadFromFile"),
			zap.Error(err))
	}

	if *SQLMode == "up" {
		_, err := organization.Exec(db, "migrate")
		if err != nil {
			log.Fatal("error creating tables",
				zap.String("package", "main"),
				zap.String("function", "organization.Exec"),
				zap.Error(err))
		}

		log.Info("UP DATABASE")
	}

	if *SQLMode == "drop" {
		_, err := organization.Exec(db, "drop")
		if err != nil {
			log.Fatal("error creating tables",
				zap.String("package", "main"),
				zap.String("function", "organization.Exec"),
				zap.Error(err))
		}
	}

	if *SQLMode == "data" {
		_, err = data.Exec(db, "data")
		if err != nil {
			log.Fatal("error creating tables",
				zap.String("package", "main"),
				zap.String("function", "data.Exec"),
				zap.Error(err))
		}
	}
}
