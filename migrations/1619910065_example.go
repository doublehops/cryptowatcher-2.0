package migrations

import (
	"cryptowatcher.example/internal/types/database"
	"errors"
	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
)

func init() {
	migorm.RegisterMigration(&migrationTest{})
}

type migrationTest struct{}

type TestTable struct {
	ID     uint32
	Name   string
	Symbol string
	Age    int32
	gorm.Model
}

func (m *migrationTest) Up(db *gorm.DB, log migorm.Logger) error {

	err := errors.New("implement me")

	db.SingularTable(true)
	db.AutoMigrate(&database.CmcHistory{})

	return err
}

func (m *migrationTest) Down(db *gorm.DB, log migorm.Logger) error {

	err := errors.New("implement me")

	db.DropTable(&TestTable{})

	return err
}
