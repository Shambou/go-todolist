package dbrepo

import (
	"fmt"
	"os"

	"github.com/Shambou/todolist/internal/config"
	"github.com/Shambou/todolist/internal/models"
	"github.com/Shambou/todolist/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *gorm.DB
}

// NewPostgresRepo creates a new postgres db repo
func NewPostgresRepo(a *config.AppConfig) repository.DatabaseRepo {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DB"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	conn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// auto migrations
	_ = conn.AutoMigrate(&models.TodoItem{})

	return &PostgresDBRepo{
		App: a,
		DB:  conn,
	}
}
