package main

import (
	"flag"

	"cryptowatcher.example/internal/funcs"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/marketmodule"
	"cryptowatcher.example/internal/pkg/orm"
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

	cmcApiKey := funcs.GetEnvironmentVar("CMC_API_KEY")
	logger := logga.New()

	db := orm.Connect(logger)

	mm := marketmodule.New(db, cmcApiKey, logger)
	_, err := mm.SaveCurrencyListing(flags.NumberToRetrieve)
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
