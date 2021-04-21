package main

import (
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"flag"
	"os"

	"cryptowatcher.example/internal/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/cmd/coinfetcher/processor"
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

	// Setup database connection.
	db := orm.Connect(logger, e)

	// Setup Coinmarketcap connection.
	cmcm := cmcmodule.New(db, e.GetVar("CMC_API_KEY"), logger)

	// Process
	//p := New(e, logger, cmcm)
	runner := processor.New(e, logger, cmcm)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	runner.Run()
	//_, err = mm.SaveCurrencyListing(flags.NumberToRetrieve)
	//if err != nil {
	//	logger.Lg.Error().Msg(err.Error())
	//}
}

func getFlags() ParamStruct {

	numberToRetrieve := flag.Int("retrieve", numberToRetrieveDefault, "Number of coins to include in fetch")
	flag.Parse()

	params := ParamStruct{
		NumberToRetrieve: *numberToRetrieve,
	}

	return params
}
