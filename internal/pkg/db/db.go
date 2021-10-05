package db

import (
	"fmt"

	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New(logger *logga.Logga, e *env.Env) (*sql.DB, error) {
	l := logger.Lg.With().Str("db", "New").Logger()

	dbName := e.GetVar("MYSQL_DATABASE")
	user := e.GetVar("MYSQL_USER")
	password := e.GetVar("MYSQL_PASSWORD")
	host := e.GetVar("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		l.Error().Msgf("unable to create db connection. %w", err)
		return db, err
	}

	return db, nil
}
