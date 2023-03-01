package main

import (
	"flag"
	"log"
	"os"

	migrate "github.com/doublehops/go-migration"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
)

func main() {
	var args migrate.Action

	flag.StringVar(&args.Action, "action", "", "the intended action")
	flag.StringVar(&args.Name, "name", "", "the name of the migration")
	flag.IntVar(&args.Number, "number", 0, "The number of migrations to run")

	configFile := flag.String("config", "config.json", "Config file to use")
	flag.Parse()

	args = setFlags(args)

	logger := logga.New()
	// Setup config.
	cfg, err := config.New(logger, *configFile)
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	args.Path = dir + "/migrations"
	args.DB = DB
	err = args.Migrate()
	if err != nil {
		os.Stderr.WriteString("There was an error initialising migration. " + err.Error() + "\n")
		os.Exit(1)
	}
}

// setFlags will check that the flags received are valid and assign default ones if not supplied.
func setFlags(args migrate.Action) migrate.Action {
	if found := args.IsValidAction(args.Action); !found {
		args.PrintHelp()
	}

	if args.Action == "create" && args.Name == "" {
		args.PrintHelp()
	}

	if args.Action == "up" && args.Number == 0 {
		args.Number = 9999 // run them all if none defined.
	}

	if args.Action == "down" && args.Number == 0 {
		args.Number = 1 // run just one if none defined.
	}

	return args
}
