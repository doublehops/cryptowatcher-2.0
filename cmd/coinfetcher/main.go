package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/doublehops/cryptowatcher-2.0/internal/aggregatorengine"
	"github.com/doublehops/cryptowatcher-2.0/internal/aggregators/coingecko"
	"github.com/doublehops/cryptowatcher-2.0/internal/aggregators/coinmarketcap"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/runflags"
)

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
		a, err = coinmarketcap.New(logger, DB, client)
	}
	// todo - remove the control statements and replace with a dynamic approach.
	if cfg.Aggregator.Name == "coingecko" {
		a, err = coingecko.New(logger, DB, client)
	}

	if a == nil {
		logger.Error("aggregator not configured.")
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
