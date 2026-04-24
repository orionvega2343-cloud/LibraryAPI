package db

import (
	"LibrariAPI/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
