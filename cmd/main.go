package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	configs "github.com/vectornko/organization-api/internal/config"
	handler "github.com/vectornko/organization-api/internal/delivery/http"
	"github.com/vectornko/organization-api/internal/repository"
	"github.com/vectornko/organization-api/internal/repository/postgres"
	"github.com/vectornko/organization-api/internal/service"
	"github.com/vectornko/organization-api/pkg/logger"
	"github.com/vectornko/organization-api/pkg/server"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// @title Organization API
// @version 1.0.0
// @description Описание API для работы с микросервисом организаций

// @host localhost:8001
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Инициализация логгера
	log := logger.NewLogger()

	// Инициализация конфига
	if err := configs.InitConfig(); err != nil {
		log.Panic("error initializing configs",
			zap.String("package", "main"),
			zap.String(" function", "initConfig"),
			zap.Error(err))
	}

	// Подключение к postgres
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Password: viper.GetString("db.password"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Panic("error initializing database",
			zap.String("package", "main"),
			zap.String("function", "postgres.NewPostgresDB"),
			zap.Error(err))
	}

	repo := repository.NewRepository(db)
	service := service.NewServices(repo)
	handler := handler.NewHandler(service)
	server := new(server.Server)

	port := viper.GetString("port")
	router := handler.InitAPI()

	// Запуск сервера
	go func() {
		err = server.Run(port, router)
		if err != nil {
			log.Info("error occured while running http server",
				zap.String("package", "main"),
				zap.String("function", "server.Run"),
				zap.Error(err))
		}
	}()
	log.Info("Start listening server...")
	log.Info("http://localhost:" + port + "/")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Server shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Error("error occured on server shutting down", zap.Error(err))
	}
}
