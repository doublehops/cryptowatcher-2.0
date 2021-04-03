package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"cryptowatcher.example/internal/pkg/marketmodule"
)

var numberToRetrieveDefault = 100
var numberToDisplayDefault = 10

type ParamStruct struct {
	NumberToRetrieve int
	NumberToDisplay  int
}

func main() {
	fmt.Println("Starting comparisons")

	flags := getFlags()

	run(flags)
}

func run(flags ParamStruct) {

	cmcApiKey := getEnvironmentVar("CMC_API_KEY")

	cmcmodule := marketmodule.New(cmcApiKey)
	_, err := cmcmodule.SaveCurrencyListing(flags.NumberToRetrieve)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getFlags() ParamStruct {

	numberToRetrieve := flag.Int("retrieve", numberToRetrieveDefault, "Number of coins to include in fetch")
	numberToDisplay := flag.Int("display", numberToDisplayDefault, "Number of coins to display in output")
	flag.Parse()

	params := ParamStruct{
		NumberToRetrieve: *numberToRetrieve,
		NumberToDisplay:  *numberToDisplay,
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
