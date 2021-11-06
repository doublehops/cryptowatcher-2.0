package main

import (
	"log"
	"os"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/runflags"
	migrate "github.com/doublehops/go-migration"
)

func main() {

	flags := runflags.GetFlags()
	logger := logga.New()
	// Setup config.
	cfg, err := config.New(logger, flags.ConfigFile)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup db connection.
	DB, err := db.New(logger, cfg.DB)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New(DB, dir+"/migrations")
	if err != nil {
		os.Stderr.WriteString("There was an error initialising migration. "+ err.Error()+"\n")
		os.Exit(1)
	}
	err = m.Migrate()
	if err != nil {
		os.Stderr.WriteString("There was an error running migration. "+ err.Error()+"\n")
		os.Exit(1)
	}
}
