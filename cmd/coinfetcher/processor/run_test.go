package processor

import (
	"cryptowatcher.example/internal/pkg/db"
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
	"cryptowatcher.example/test"
)

var l *logga.Logga
var e *env.Env
var DB *sql.DB
var tx *sql.Tx

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	e, err := env.New(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	l = logga.New()
	DB, err = db.New(l, e)
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

	chm := cmcmodule.New(e.GetVar("CMC_API_KEY"), server.URL, l)

	p := New(e, l, tx, chm)
	p.Run()

	var currencies cmcmodule.Data
	json.Unmarshal(testJsonResponse, &currencies)

	jsonRec1 := currencies.Currencies[0]

	// Test record in currency table.
	var curDbRec1 database.Currency

	cm := currency.New(tx, l)
	cm.GetRecordBySymbol(&curDbRec1, jsonRec1.Symbol)

	assert.Equal(t, jsonRec1.Name, curDbRec1.Name)
	assert.Equal(t, jsonRec1.Symbol, curDbRec1.Symbol)
}
