package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/carprice-tech/migorm"
	"errors"
)

func init(){
	migorm.RegisterMigration(&migrationAnotherOne{})
}

type migrationAnotherOne struct{}

func (m *migrationAnotherOne) Up(db *gorm.DB, log migorm.Logger) error{

	err := errors.New("implement me")

	return err
}

func (m *migrationAnotherOne) Down(db *gorm.DB, log migorm.Logger) error{

	err := errors.New("implement me")

	return err
}
