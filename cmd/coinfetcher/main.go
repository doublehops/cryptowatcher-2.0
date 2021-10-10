package main

import (
	"os"

	"cryptowatcher.example/cmd/coinfetcher/processor"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/runflags"
)

var numberToRetrieveDefault = 10 // @todo - this var can be removed or better handled elsewhere.

func main() {
	flags := runflags.GetFlags()
	run(flags)
}

func run(flags runflags.FlagStruct) {

	// Setup logger.
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

	// Setup Coinmarketcap connection.
	cmcm := cmcmodule.New(cfg.Tracker.Host, cfg.Tracker.APIKey, logger)

	// Process
	runner := processor.New(cfg.Tracker, logger, DB, cmcm)
	err = runner.Run()
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}
