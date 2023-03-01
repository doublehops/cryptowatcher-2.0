package coingecko

import (
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"github.com/doublehops/cryptowatcher-2.0/internal/aggregatorengine"
	"github.com/doublehops/cryptowatcher-2.0/internal/models/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
	"github.com/doublehops/cryptowatcher-2.0/test/testfuncs"
)

var (
	l   *logga.Logga
	cfg *config.Config
	DB  *sql.DB
	tx  *sql.Tx
)

func setup(t *testing.T) {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	var err error

	// Setup aggregatorConfig.
	cfg, err = testfuncs.GetTestConfig(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		t.Errorf("unable to get config. %s", err)
	}

	l = logga.New()
	DB, err = db.New(l, cfg.DB)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		t.Errorf("unable to create database connection. %s", err)
	}
	tx, err = DB.Begin()
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		t.Errorf("unable to begin database transaction. %s", err)
	}
}

func tearDown(t *testing.T) {
	err := tx.Rollback()
	if err != nil {
		t.Errorf("unable to rollback transcaction. %s", err)
	}
}

func TestRun(t *testing.T) {
	setup(t)
	defer tearDown(t)

	// Setup dbConn connection.
	dbConn, err := db.New(l, cfg.DB)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	// Setup test http server.
	testJSONResponse, err := testfuncs.GetTestJSONResponse("coingecko_coin_list_response.json")
	if err != nil {
		t.Fatalf("error getting server response. %s", err)
	}
	server := testfuncs.SetupTestServer(testJSONResponse)
	defer server.Close()

	aggConfig := &aggregatorConfig{
		Name:  "CoinGecko-test",
		Label: "coingecko-test",
		HostConfig: HostConfig{
			APIHost: server.URL,
		},
	}

	runner := &Runner{
		aggregatorConfig: aggConfig,
		l:                l,
		db:               dbConn,
		client:           server.Client(),
	}

	agg := aggregatorengine.New(DB, runner, l)
	err = agg.UpdateLatestHistory()
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	var currencies []*Currency
	err = json.Unmarshal(testJSONResponse, &currencies)
	if err != nil {
		t.Errorf("could not unmarshal JSON. %s", err)
	}

	jsonRec1 := currencies[0]

	// Test record in currency table.
	var curDBRec1 database.Currency

	cm := currency.New(DB, l)
	err = cm.GetRecordBySymbol(&curDBRec1, jsonRec1.Symbol)
	if err != nil {
		t.Errorf("error with GetRecordBySymbol. %s", err)
	}

	if jsonRec1.Name != curDBRec1.Name {
		t.Errorf("name not as expected. Got: %s; found: %s;", jsonRec1.Name, curDBRec1.Name)
	}

	if jsonRec1.Symbol != curDBRec1.Symbol {
		t.Errorf("symbol not as expected. Got: %s; found: %s;", jsonRec1.Symbol, curDBRec1.Symbol)
	}

	err = cm.DeleteRecord(curDBRec1.ID)
	if err != nil {
		t.Errorf("Unable to remove test record from database")
	}
}
