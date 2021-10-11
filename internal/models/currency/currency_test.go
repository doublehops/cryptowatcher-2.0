package currency

import (
	"database/sql"
	"os"
	"testing"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

var l *logga.Logga
var DB *sql.DB
var tx *sql.Tx

var testCoin *sql.DB

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	// Setup logger.
	l = logga.New()

	// Setup config.
	cfg, err := config.New(l, "../../../config.json.test")
	if err != nil {
		l.Lg.Error().Msgf("error starting main. %w", err.Error())
		os.Exit(1)
	}
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

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

	record := getTestRecord()
	cm := New(tx, l)

	_, err = cm.CreateRecord(record)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
}

func teardown() {
	tx.Rollback()
}

func TestCreateRecord(t *testing.T) {

	setup()
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

	setup()
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
