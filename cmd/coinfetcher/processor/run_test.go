package processor

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"cryptowatcher.example/internal/models/cmchistory"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
)

var l *logga.Logga
var db *gorm.DB
var tx *gorm.DB
var e *env.Env

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	e, err := env.New(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	l = logga.New()
	db = orm.Connect(l, e)
	tx = db.Begin()
}

func teardown() {
	tx.Rollback()
}

func TestRun(t *testing.T) {

	setup()
	defer teardown()

	absPath, _ := filepath.Abs("./../../../dev/test_cmc_list_response.json")
	testJsonResponse, _ := ioutil.ReadFile(absPath)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(testJsonResponse))
	}))
	defer server.Close()

	baseUrl := server.URL

	cmcm := cmcmodule.New(e.GetVar("CMC_API_KEY"), baseUrl, l)

	p := New(e, l, tx, cmcm)
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

	// Test records in cmc_history table.
	cmm := cmchistory.New(tx, l)
	records, err := cmm.GetRecordsBySymbol(jsonRec1.Symbol)
	assert.Nil(t, err, "Getting cmchistory records returned no error")

	assert.Equal(t, 1, len(records))
	r1 := records[0]
	fmt.Printf("history records found: %d\n", len(records))
	assert.Equal(t, jsonRec1.Name, r1.Name)
	assert.Equal(t, jsonRec1.Symbol, r1.Symbol)
	assert.Equal(t, jsonRec1.NumMarketPairs, r1.NumMarketPairs)
	assert.Equal(t, jsonRec1.Slug, r1.Slug)
	assert.Equal(t, jsonRec1.NumMarketPairs, r1.NumMarketPairs)
	assert.Equal(t, jsonRec1.DateAdded, r1.DateAdded)
	assert.Equal(t, jsonRec1.MaxSupply, r1.MaxSupply)
	assert.Equal(t, jsonRec1.CirculatingSupply, r1.CirculatingSupply)
	assert.Equal(t, jsonRec1.TotalSupply, r1.TotalSupply)
	assert.Equal(t, jsonRec1.CmcRank, r1.CmcRank)
	assert.Equal(t, jsonRec1.CirculatingSupply, r1.CirculatingSupply)
}
