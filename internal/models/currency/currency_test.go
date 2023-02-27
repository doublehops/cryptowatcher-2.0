package currency

import (
	"database/sql"
	"github.com/doublehops/cryptowatcher-2.0/test/testfuncs"
	"os"
	"testing"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

var l *logga.Logga
var DB *sql.DB
var tx *sql.Tx

var testCoin *sql.DB

func setup(t *testing.T) {
	_ = os.Setenv("APP_ENV", "test")

	// Setup logger.
	l = logga.New()

	// Setup config.
	cfg, err := testfuncs.GetTestConfig(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		t.Errorf("unable to get config. %s", err)
	}

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

	record := getTestRecord()
	cm := New(tx, l)

	_, err = cm.CreateRecord(record)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		t.Errorf("unable to create record. %s", err)
	}
}

func teardown() {
	tx.Rollback()
}

func TestCreateRecord(t *testing.T) {

	setup(t)
	defer teardown()

	cm := New(tx, l)

	cr := &database.Currency{
		Name:   "createTestCoin",
		Symbol: "TestCoin",
	}

	ID, err := cm.CreateRecord(cr)
	if err != nil {
		t.Errorf("error getting record by ID. %s", err)
	}
	if ID == 0 {
		t.Errorf("Record creation did not return a last insert ID")
	}

	var rr database.Currency
	err = cm.GetRecordByID(&rr, ID)
	if err != nil {
		t.Errorf("error getting record by ID. %s", err)
	}
	if int64(rr.ID) != ID {
		t.Errorf("record id not as expected. Got: %d; wanted: %d;", cr.ID, ID)
	}
	if rr.Name != cr.Name {
		t.Errorf("record id not as expected. Got: %s; wanted: %s;", rr.Name, cr.Name)
	}
}

func TestGetRecord(t *testing.T) {

	setup(t)
	defer teardown()

	cm := New(tx, l)

	record := getTestRecord()

	var tc database.Currency

	err := cm.GetRecordBySymbol(&tc, record.Symbol)
	if err != nil {
		t.Fatalf("error getting record. %s", err)
	}
	if record.Name != tc.Name {
		t.Fatalf("Name not as expected. Got: %s; wanted: %s;", record.Name, tc.Name)
	}
}

func getTestRecord() *database.Currency {

	return &database.Currency{
		Name:   "testcoin",
		Symbol: "TestCoin",
	}
}
