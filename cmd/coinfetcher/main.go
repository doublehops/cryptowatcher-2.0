package main

import (
	"flag"
	"os"

	"cryptowatcher.example/cmd/coinfetcher/processor"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
)

var numberToRetrieveDefault = 10 // @todo - this var can be removed or better handled elsewhere.

type ParamStruct struct {
	NumberToRetrieve int
}

func main() {
	flags := getFlags()
	run(flags)
}

func run(flags ParamStruct) {

	RequiredEnvVars := []string{
		"CMC_API_KEY",
		"CMC_API_HOST",

		"MYSQL_HOST",
		"MYSQL_USER",
		"MYSQL_PASSWORD",
		"MYSQL_DATABASE",
	}

	// Setup logger.
	logger := logga.New()

	// Setup environment.
	e, err := env.New(logger)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	err = e.TestEnvironmentVars(RequiredEnvVars)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup db connection.
	db, err := db.New(logger, e)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup Coinmarketcap connection.
	cmcm := cmcmodule.New(e.GetVar("CMC_API_KEY"), e.GetVar("CMC_API_HOST"), logger)

	// Process
	runner := processor.New(e, logger, db, cmcm)
	err = runner.Run()
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}

func getFlags() ParamStruct {

	numberToRetrieve := flag.Int("retrieve", numberToRetrieveDefault, "Number of coins to include in fetch")
	flag.Parse()

	params := ParamStruct{
		NumberToRetrieve: *numberToRetrieve,
	}

	return params
}
