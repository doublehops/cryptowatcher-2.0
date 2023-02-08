package main

import (
	"fmt"
	"os"

	"cryptowatcher.example/internal/aggregatorengine"
	"cryptowatcher.example/internal/aggregators/coinmarketcap"
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
	//cmcm := coinmarketcap.New(cfg.Aggregator, logger)

	//agggg, _ := coinmarketcap.New(cfg.Aggregator, nil, nil)

	// Process
	//cmc := coinmarketcap.New(cfg.Aggregator, logger, DB, cmcm)

	var aggg aggregatorengine.Aggregator
	// todo - remove the control statements and replace with a dynamic approach.
	if cfg.Aggregator.Name == "coinmarketcap" {
		aggg, err = coinmarketcap.New(logger, DB)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("unable to instantiate aggregator. %s", err))
		os.Exit(1)
	}

	agg := aggregatorengine.New(DB, logger)
	err = agg.UpdateLatestHistory(aggg)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}
