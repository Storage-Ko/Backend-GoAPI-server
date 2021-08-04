package db

import (
	"fmt"

	"github.com/Backend-GoAPI-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func Start() (*gorm.DB, error) {
	var dbConfig map[string]string
	dbConfig, err := godotenv.Read()
	if err != nil {
		utils.HandleErr(err)
		return nil, err
	}

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig["MYSQL_USER"],
		dbConfig["MYSQL_PASSWORD"],
		dbConfig["MYSQL_PROTOCOL"],
		dbConfig["MYSQL_HOST"],
		dbConfig["MYSQL_PORT"],
		dbConfig["MYSQL_DBNAME"],
	)

	db, err := gorm.Open("mysql", mysqlCredentials)
	return db, err
}
