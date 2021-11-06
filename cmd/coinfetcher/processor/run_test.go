package processor

import (
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

var l *logga.Logga
var cfg *config.Config
var DB *sql.DB
var tx *sql.Tx

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	// Setup config.
	var err error
	cfg, err = config.New(l, "../../../config.json.test")
	if err != nil {
		l.Lg.Error().Msgf("error starting main. %w", err.Error())
		os.Exit(1)
	}

	l = logga.New()
	DB, err = db.New(l, cfg.DB)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
	tx, err = DB.Begin()
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}

func tearDown() {
	tx.Rollback()
}

func TestRun(t *testing.T) {

	setup()
	defer tearDown()

	// Setup test http server.
	testJsonResponse, err := testfuncs.GetServerResponse("test_cmc_list_response.json")
	if err != nil {
		t.Fatalf("error getting server response. %s", err)
	}
	server := testfuncs.SetupTestServer(testJsonResponse)
	defer server.Close()

	cfg.Tracker.Host = server.URL // Set URL to that of the test response

	chm := cmcmodule.New(cfg.Tracker, l)

	p := New(cfg.Tracker, l, DB, chm)
	err = p.Run()
	if err != nil {
		t.Errorf("unable to instantiate runner. %s", err)
	}

	var currencies cmcmodule.Data
	err = json.Unmarshal(testJsonResponse, &currencies)
	if err != nil {
		t.Errorf("could not unmarshal JSON. %s", err)
	}

	jsonRec1 := currencies.Currencies[0]

	// Test record in currency table.
	var curDbRec1 database.Currency

	cm := currency.New(DB, l)
	err = cm.GetRecordBySymbol(&curDbRec1, jsonRec1.Symbol)
	if err != nil {
		t.Errorf("error with GetRecordBySymbol. %s", err)
	}

	if jsonRec1.Name != curDbRec1.Name {
		t.Errorf("name not as expected. Got: %s; found: %s;", jsonRec1.Name, curDbRec1.Name)
	}

	if jsonRec1.Name != curDbRec1.Name {
		t.Errorf("symbol not as expected. Got: %s; found: %s;", jsonRec1.Symbol, curDbRec1.Symbol)
	}
}
