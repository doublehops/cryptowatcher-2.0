package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/marketmodule"
)

var numberToRetrieveDefault = 100

type ParamStruct struct {
	NumberToRetrieve int
}

func main() {
	flags := getFlags()

	run(flags)
}

func run(flags ParamStruct) {

	cmcApiKey := getEnvironmentVar("CMC_API_KEY")
	logga := logga.New()

	mm := marketmodule.New(cmcApiKey, logga)
	_, err := mm.SaveCurrencyListing(flags.NumberToRetrieve)
	if err != nil {
		logga.Lg.Error().Msg(err.Error())
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

func getEnvironmentVar(varName string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to open environment file")
		os.Exit(1)
	}

	return os.Getenv(varName)
}
