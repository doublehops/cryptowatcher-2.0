package main

import (
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
	"os"
)

func main() {

	// Setup logger.
	logger := logga.New()

	// Setup environment.
	e, err := env.New(logger)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	db := orm.Connect(logger, e)

	db.Migrator().AutoMigrate(&database.Currency{})

	db.Migrator().AutoMigrate(&database.CmcHistory{})
}
