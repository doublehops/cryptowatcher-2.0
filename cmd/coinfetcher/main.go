package main

import (
	"fmt"
	"net/http"
	"os"

	"cryptowatcher.example/internal/aggregatorengine"
	"cryptowatcher.example/internal/aggregators/coingecko"
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
	
	client := &http.Client{}

	var a aggregatorengine.Aggregator
	// todo - remove the control statements and replace with a dynamic approach.
	if cfg.Aggregator.Name == "coinmarketcap" {
		a, err = coinmarketcap.New(logger, DB, client, flags.ConfigFile)
	}
	// todo - remove the control statements and replace with a dynamic approach.
	if cfg.Aggregator.Name == "coingecko" {
		a, err = coingecko.New(logger, DB)
	}

	if a == nil {
		logger.Error(fmt.Sprintf("aggregator not configured."))
		os.Exit(1)
	}

	if err != nil {
		logger.Error(fmt.Sprintf("unable to instantiate aggregator. %s", err))
		os.Exit(1)
	}

	agg := aggregatorengine.New(DB, a, logger)
	err = agg.UpdateLatestHistory()
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}
