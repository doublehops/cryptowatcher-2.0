package main

import (
	"flag"
	"os"
	
	"cryptowatcher.example/internal/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/marketmodule"
	"cryptowatcher.example/internal/pkg/orm"
	envtype "cryptowatcher.example/internal/types/env"
)

var numberToRetrieveDefault = 10

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

	logger := logga.New()

	envVars := envtype.EnvVars{}
	err := env.GetEnvironmentVars(&envVars, RequiredEnvVars)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	db := orm.Connect(logger, envVars)

	mm := marketmodule.New(db, envVars["CMC_API_KEY"], logger)
	_, err = mm.SaveCurrencyListing(flags.NumberToRetrieve)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
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
