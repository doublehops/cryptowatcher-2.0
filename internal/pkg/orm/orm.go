package orm

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
)

func Connect(logger *logga.Logga, e *env.Env) *gorm.DB {

	l := logger.Lg.With().Str("marketmodule", "GetCurrencyListing").Logger()

	dbName := e.GetVar("MYSQL_DATABASE")
	user := e.GetVar("MYSQL_USER")
	password := e.GetVar("MYSQL_PASSWORD")
	host := e.GetVar("MYSQL_HOST")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName), // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		l.Error().Msg("Error establishing database connection")
		l.Error().Msgf("%v", err)
		os.Exit(1)
	}

	return db
}
