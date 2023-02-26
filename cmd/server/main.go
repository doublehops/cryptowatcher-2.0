package main

import (
	"os"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/router"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/runflags"
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
