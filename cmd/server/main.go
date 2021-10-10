package main

import (
	"os"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/router"
	"cryptowatcher.example/internal/pkg/runflags"
	"github.com/gin-gonic/gin"
)

func main() {

	flags := runflags.GetFlags()

	// Setup logger.
	l := logga.New()

	// Setup config.
	cfg, err := config.New(l, flags.ConfigFile)
	if err != nil {
		l.Lg.Error().Msgf("error starting main. %w", err.Error())
		os.Exit(1)
	}

	// Setup db connection.
	DB, err := db.New(l, cfg.DB)
	if err != nil {
		l.Lg.Error().Msgf("error creating database connection. %w", err.Error())
		os.Exit(1)
	}

	r := gin.Default()
	router.New(r, DB, l)

	r.Run(":8080")
}
