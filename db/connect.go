package db

import (
	"fmt"

	"github.com/Backend-GoAPI-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func Start() {
	var dbConfig map[string]string
	dbConfig, err := godotenv.Read()
	utils.HandleErr(err)

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig["MYSQL_USER"],
		dbConfig["MYSQL_PASSWORD"],
		dbConfig["MYSQL_PROTOCOL"],
		dbConfig["MYSQL_HOST"],
		dbConfig["MYSQL_PORT"],
		dbConfig["MYSQL_DBNAME"],
	)

	db, err = gorm.Open("mysql", mysqlCredentials)
	utils.HandlePanic(err)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
