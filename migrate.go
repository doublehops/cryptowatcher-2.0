package main

import (
	"flag"
	//"fmt"
	"os"

	//"github.com/carprice-tech/migorm"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"

	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
)

type ParamStruct struct {
	env string
}

func main() {

	flags := getFlags()

	if flags.env != "ignore" {
		_ = os.Setenv("APP_ENV", flags.env)
	}

	// Setup logger.
	logger := logga.New()

	// Setup environment.
	e, err := env.New(logger)
	if err != nil {
		logger.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}

	/**
	 * Frustratingly I am using Gorm version2 (gorm.io/gorm) but the Gorm migration tool only works with
	 * Gorm 1.x (github.com/jinzhu/gorm). Therefore the main application is running Gorm version 2 but for now the
	 * migration tool in this file is using Gorm 1.x. I hope they're compatible.
	 */

	/**
	 * This is to make a database connection using Gorm 2.0 (gorm.io/gorm)
	 *
	 * db := orm.Connect(logger, e)
	 */
	db := orm.Connect(logger, e)

	// Version 1.x connection
	//db, err := getVersionOneDatabaseConn(e)
	//if err != nil{
	//	panic(err)
	//}
	//
	//migrater := migorm.NewMigrater(db)
	//migorm.Run(migrater)

	db.Migrator().DropTable(&database.Currency{})
	db.Migrator().AutoMigrate(&database.Currency{})

	db.Migrator().DropTable(&database.CmcHistory{})
	db.Migrator().AutoMigrate(&database.CmcHistory{})
}

//func getVersionOneDatabaseConn(e *env.Env) (*gorm.DB, error){
//
//	user := e.GetVar("MYSQL_USER")
//	pass := e.GetVar("MYSQL_PASSWORD")
//	host := e.GetVar("MYSQL_HOST")
//	name := e.GetVar("MYSQL_DATABASE")
//
//	conStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=true&loc=Local", user, pass, host, name)
//	dbConn, err := gorm.Open("mysql", conStr)
//
//	return dbConn, err
//}

func getFlags() ParamStruct {

	e := flag.String("env", "ignore", "Which environment to use")
	flag.Parse()

	params := ParamStruct{
		env: *e,
	}

	return params
}
