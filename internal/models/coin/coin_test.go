package coin

import (
	"gorm.io/gorm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

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

	// Add test record
	cm := New(tx, l)
	testCoin = cm.CreateCoin(getTestRecord())

	// call flag.Parse() here if TestMain uses flags
	//os.Exit(m.Run())
}

func teardown() {
	tx.Rollback()
}

func TestCreateRecord(t *testing.T) {

	setup()
	defer teardown()

	cm := New(tx, l)

	cr := &database.Coin{
		Name: "createTestCoin",
		Symbol: "CTestCoin",
	}

	result:= cm.CreateCoin(cr)
	if result.Error != nil {
		t.Errorf("Error creating coin record")
	}
	assert.Nil(t, result.Error, "Record created without error")
}

func TestGetRecord(t *testing.T) {

	setup()
	defer teardown()

	cm := New(tx, l)

	tcoin := getTestRecord()

	tc := cm.GetCoinBySymbol(tcoin.Symbol)
	if tc.Name != tcoin.Name {
		t.Errorf("Error getting coin record")
	}
}

func getTestRecord() *database.Coin {

	return &database.Coin{
		Name: "testcoin",
		Symbol: "TestCoin",
	}
}
