package cmchistory

import (
	"os"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
)

var l *logga.Logga
var db *gorm.DB
var tx *gorm.DB

var testCoin *gorm.DB

func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()
	db = orm.Connect(l)
	tx = db.Begin()
}

func teardown() {
	tx.Rollback()
}

func TestCreateAndRetrieveRecord(t *testing.T) {

	setup()
	defer teardown()

	cmcm := New(tx, l)

	r := &database.CmcHistory{
		Name:              fake.CharactersN(10),
		Symbol:            fake.Characters(),
		Slug:              fake.Characters(),
		NumMarketPairs:    12,
		DateAdded:         "2021-04-10 10:46:00",
		MaxSupply:         12.32,
		CirculatingSupply: 12313123123.02,
		TotalSupply:       564654.40,
		CmcRank:           123,
		QuotePrice:        100.23,
		Volume24h:         154867894541.151,
		PercentChange1h:   1,
		PercentChange24h:  24,
		PercentChange7D:   7,
		PercentChange30D:  30,
		PercentChange60D:  60,
		PercentChange90D:  90,
		MarketCap:         1222333555,
	}

	record, err := cmcm.CreateRecord(r)

	assert.Nil(t, err, "Created record returned no error")

	rr, err := cmcm.GetRecordByID(record.ID)

	assert.Equal(t, r.Name, rr.Name, "Name returned as expected")
	assert.Equal(t, r.Symbol, rr.Symbol, "Name returned as expected")
	assert.Equal(t, r.TotalSupply, rr.TotalSupply, "Name returned as expected")
}
