package main

import (
	"fmt"

	"github.com/Backend-GoAtreugo-server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
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

	db, err := gorm.Open("mysql", mysqlCredentials)
	defer db.Close()

	utils.HandleErr(err)

}
