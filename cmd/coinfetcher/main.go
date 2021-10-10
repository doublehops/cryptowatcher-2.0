package main

import (
	"flag"
	"os"

	"cryptowatcher.example/cmd/coinfetcher/processor"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
)

var numberToRetrieveDefault = 10 // @todo - this var can be removed or better handled elsewhere.

type FlagStruct struct {
	ConfigFile string
}

func main() {
	flags := getFlags()
	run(flags)
}

func run(flags FlagStruct) {

	// Setup logger.
	logger := logga.New()

	// Setup environment.
	cfg, err := config.New(logger, flags.ConfigFile)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup db connection.
	db, err := db.New(logger, cfg.DB)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup Coinmarketcap connection.
	cmcm := cmcmodule.New(cfg.Tracker.Host, cfg.Tracker.APIKey, logger)

	// Process
	runner := processor.New(cfg.Tracker, logger, db, cmcm)
	err = runner.Run()
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}

func getFlags() FlagStruct {

	configFile := flag.String("config", "config.json", "Config file to use")
	flag.Parse()

	params := FlagStruct{
		ConfigFile: *configFile,
	}

	return params
}
