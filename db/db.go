package db

import (
	"database/sql"
	"fmt"

	"github.com/Backend-GoAtreugo-server/model"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() *sql.DB {
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

	db, err := gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})
	utils.HandleErr(err)

	mysql, err := db.DB()
	utils.HandleErr(err)
	mysql.Begin()
	db.AutoMigrate(&model.Admin{}, &model.User{}, &model.Comment{})

	return mysql
}
