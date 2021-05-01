package main

import (
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
)

type ParamStruct struct {
	env string
}

func main() {

	flags := getFlags()
	fmt.Println(flags.env)
	spew.Dump(flags.env)

	if flags.env != "ignore" {
		_ = os.Setenv("APP_ENV", flags.env)
	}

	// Setup logger.
	logger := logga.New()

	// Setup environment.
	e, err := env.New(logger)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	db := orm.Connect(logger, e)

	db.Migrator().DropTable(&database.Currency{})
	db.Migrator().AutoMigrate(&database.Currency{})

	db.Migrator().DropTable(&database.CmcHistory{})
	db.Migrator().AutoMigrate(&database.CmcHistory{})
}

func getFlags() ParamStruct {

	env := flag.String("env", "ignore", "Which environment to use")
	flag.Parse()

	params := ParamStruct{
		env: *env,
	}

	return params
}
