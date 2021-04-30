package currency

import (
	"os"
	"testing"

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

var testCoin *gorm.DB


func setup() {
	_ = os.Setenv("APP_ENV", "test")

	l = logga.New()

	e, err := env.New(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
	db = orm.Connect(l, e)
	tx = db.Begin()

	// Add test record
	record := getTestRecord()
	cm := New(tx, l)
	err = cm.CreateCurrency(record)
}

func teardown() {
	tx.Rollback()
}

func TestCreateRecord(t *testing.T) {

	setup()
	defer teardown()

	cm := New(tx, l)

	cr := &database.Currency{
		Name: "createTestCoin",
		Symbol: "CTestCoin",
	}

	err := cm.CreateCurrency(cr)
	assert.Nil(t, err, "Record created without error")
}

func TestGetRecord(t *testing.T) {

	setup()
	defer teardown()

	cm := New(tx, l)

	cur := getTestRecord()

	var tc database.Currency

	cm.GetCurrencyBySymbol(&tc, cur.Symbol)
	assert.Equal(t, cur.Name, tc.Name, "Retrieved currency")
}

func getTestRecord() *database.Currency {

	return &database.Currency{
		Name: "testcoin",
		Symbol: "TestCoin",
	}
}
