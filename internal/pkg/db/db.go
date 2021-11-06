package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/logga"
)

func New(logger *logga.Logga, cfg config.DB) (*sql.DB, error) {
	l := logger.Lg.With().Str("db", "New").Logger()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Pass, cfg.Host, cfg.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		l.Error().Msgf("unable to create db connection. %w", err)
		return db, err
	}

	return db, nil
}
