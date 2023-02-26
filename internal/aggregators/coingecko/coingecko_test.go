package coingecko

import (
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"cryptowatcher.example/internal/aggregatorengine"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
	"cryptowatcher.example/test/testfuncs"
)

var l *logga.Logga
var cfg *config.Config
var DB *sql.DB
var tx *sql.Tx

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	// Setup aggregatorConfig.
	var err error
	cfg, err = config.New(l, "../../../config.json.test")
	if err != nil {
		l.Lg.Error().Msgf("error starting main. %s", err.Error())
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
	testJsonResponse, err := testfuncs.GetTestJsonResponse("coingecko_coin_list_response.json")
	if err != nil {
		t.Fatalf("error getting server response. %s", err)
	}
	server := testfuncs.SetupTestServer(testJsonResponse)
	defer server.Close()

	//cfg.Aggregator = server.URL // Set URL to that of the test response

	//runner := New(cfg.Aggregator, l)

	coingecko, err := New(l, DB)
	if err != nil {
		t.Errorf("error instantiating aggregator. %s", err)
	}

	agg := aggregatorengine.New(DB, coingecko, l)
	err = agg.UpdateLatestHistory()
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	var currencies CurrencyData
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

	if jsonRec1.Symbol != curDbRec1.Symbol {
		t.Errorf("symbol not as expected. Got: %s; found: %s;", jsonRec1.Symbol, curDbRec1.Symbol)
	}

	err = cm.DeleteRecord(curDbRec1.ID)
	if err != nil {
		t.Errorf("Unable to remove test record from database")
	}

}
