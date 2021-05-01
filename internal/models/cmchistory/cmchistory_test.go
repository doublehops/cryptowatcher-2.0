package cmchistory

import (
	"cryptowatcher.example/internal/models/currency"
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
)

var l *logga.Logga
var db *gorm.DB
var tx *gorm.DB

var cr *database.Currency

func setup() {

	_ = os.Setenv("APP_ENV", "test")

	e, err := env.New(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	l = logga.New()

	db = orm.Connect(l, e)
	tx = db.Begin()
	cr = createTestRecords(l)
}

func teardown() {
	tx.Rollback()
}

func createTestRecords(l *logga.Logga) *database.Currency {

	lg := l.Lg.With().Str("cmchistory_test", "createTestRecords").Logger()

	cm := currency.New(tx, l)

	cr := database.Currency{
		Name:   fake.CharactersN(5),
		Symbol: fake.CharactersN(3),
	}

	err := cm.CreateCurrency(&cr)
	if err != nil {
		lg.Error().Msgf("Unable to create test record.")
	}

	return &cr
}

func TestCreateAndRetrieveRecord(t *testing.T) {

	setup()
	defer teardown()

	r := &database.CmcHistory{
		CurrencyID:        cr.ID,
		Name:              fake.CharactersN(10),
		Symbol:            fake.Characters(),
		Slug:              fake.Characters(),
		NumMarketPairs:    12,
		DateAdded:         "2021-04-10 10:46:00",
		MaxSupply:         12.32,
		CirculatingSupply: 123133.02,
		TotalSupply:       564654.40,
		CmcRank:           123,
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

	err := chm.CreateRecord(r)
	assert.Nil(t, err, "Created record returned no error")

	var rt database.CmcHistory

	err = chm.GetRecordByID(&rt, r.ID)
	assert.Nil(t, err, "Get record returned no error")

	assert.Equal(t, cr.ID, rt.CurrencyID, "Record returned as expected")
	assert.Equal(t, r.Name, rt.Name, "Record returned as expected")
	assert.Equal(t, r.Symbol, rt.Symbol, "Record returned as expected")
	assert.Equal(t, r.Slug, rt.Slug, "Record returned as expected")
	assert.Equal(t, r.NumMarketPairs, rt.NumMarketPairs, "Record returned as expected")
	assert.Equal(t, r.MaxSupply, rt.MaxSupply, "Record returned as expected")
	assert.Equal(t, r.CirculatingSupply, rt.CirculatingSupply, "Record returned as expected")
	assert.Equal(t, r.TotalSupply, rt.TotalSupply, "Record returned as expected")
	assert.Equal(t, r.CmcRank, rt.CmcRank, "Record returned as expected")
	assert.Equal(t, r.QuotePrice, rt.QuotePrice, "Record returned as expected")
	assert.Equal(t, r.Volume24h, rt.Volume24h, "Record returned as expected")
	assert.Equal(t, r.PercentChange1h, rt.PercentChange1h, "Record returned as expected")
	assert.Equal(t, r.PercentChange24h, rt.PercentChange24h, "Record returned as expected")
	assert.Equal(t, r.PercentChange7D, rt.PercentChange7D, "Record returned as expected")
	assert.Equal(t, r.PercentChange30D, rt.PercentChange30D, "Record returned as expected")
	assert.Equal(t, r.PercentChange60D, rt.PercentChange60D, "Record returned as expected")
	assert.Equal(t, r.PercentChange90D, rt.PercentChange90D, "Record returned as expected")
	assert.Equal(t, r.MarketCap, rt.MarketCap, "Record returned as expected")
}
