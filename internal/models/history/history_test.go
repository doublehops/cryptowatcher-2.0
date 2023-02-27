package history

import (
	"database/sql"
	"os"
	"testing"

	"github.com/icrowley/fake"

	"github.com/doublehops/cryptowatcher-2.0/internal/models/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/db"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

var l *logga.Logga
var DB *sql.DB
var tx *sql.Tx

var cr *database.Currency

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
	createTestRecords(l)
}

func teardown() {
	tx.Rollback()
}

func createTestRecords(l *logga.Logga) {

	lg := l.Lg.With().Str("history_test", "createTestRecords").Logger()

	cm := currency.New(tx, l)

	cr = &database.Currency{
		Name:   fake.CharactersN(5),
		Symbol: fake.CharactersN(3),
	}

	lastInsertID, err := cm.CreateRecord(cr)
	if err != nil {
		lg.Error().Msgf("Unable to create test record.")
	}

	lg.Debug().Msgf("added history record with ID: %d", lastInsertID)
}

func TestCreateAndRetrieveRecord(t *testing.T) {

	setup()
	defer teardown()

	r := &database.History{
		AggregatorID:      1,
		Name:              fake.CharactersN(10),
		Symbol:            fake.Characters(),
		Slug:              fake.Characters(),
		NumMarketPairs:    12,
		DateAdded:         "2021-04-10 10:46:00",
		MaxSupply:         12.32,
		CirculatingSupply: 123133.02,
		TotalSupply:       564654.40,
		Rank:              123,
		QuotePrice:        100.23,
		Volume24h:         15541.15,
		PercentChange1h:   1,
		PercentChange24h:  24,
		PercentChange7D:   7,
		PercentChange30D:  30,
		PercentChange60D:  60,
		PercentChange90D:  90,
		MarketCap:         12555,
	}

	chm := New(tx, l)

	var err error
	r.ID, err = chm.CreateRecord(r)
	if err != nil {
		t.Errorf("could not create record. %s", err)
	}

	var rt database.History

	err = chm.GetRecordByID(&rt, r.ID)
	if err != nil {
		t.Errorf("Get record returned no error. %s", err)
	}

	tests := []struct {
		Name     string
		Got      interface{}
		Expected interface{}
	}{
		{Name: "Currency ID", Got: rt.Name, Expected: r.Name},
		{Name: "Symbol", Got: rt.Symbol, Expected: r.Symbol},
		{Name: "Slug", Got: rt.Slug, Expected: r.Slug},
		{Name: "NumMarketPairs", Got: rt.NumMarketPairs, Expected: r.NumMarketPairs},
		{Name: "MaxSupply", Got: rt.MaxSupply, Expected: r.MaxSupply},
		{Name: "CirculatingSupply", Got: rt.CirculatingSupply, Expected: r.CirculatingSupply},
		{Name: "TotalSupply", Got: rt.TotalSupply, Expected: r.TotalSupply},
		{Name: "CmcRank", Got: rt.Rank, Expected: r.Rank},
		{Name: "QuotePrice", Got: rt.QuotePrice, Expected: r.QuotePrice},
		{Name: "Volume24h", Got: rt.Volume24h, Expected: r.Volume24h},
		{Name: "PercentChange1h", Got: rt.PercentChange1h, Expected: r.PercentChange1h},
		{Name: "PercentChange24h", Got: rt.PercentChange24h, Expected: r.PercentChange24h},
		{Name: "PercentChange7D", Got: rt.PercentChange7D, Expected: r.PercentChange7D},
		{Name: "PercentChange30D", Got: rt.PercentChange30D, Expected: r.PercentChange30D},
		{Name: "PercentChange60D", Got: rt.PercentChange60D, Expected: r.PercentChange60D},
		{Name: "PercentChange90D", Got: rt.PercentChange90D, Expected: r.PercentChange90D},
		{Name: "MarketCap", Got: rt.MarketCap, Expected: r.MarketCap},
	}

	for _, tt := range tests {
		if tt.Got != tt.Expected {
			t.Fatalf("Test (%s) - Values do not match. Got: %v; Expected: %v", tt.Name, tt.Got, tt.Expected)
		}
	}
}
