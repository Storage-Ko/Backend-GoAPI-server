package db

import (
	"fmt"

	"github.com/Backend-GoAtreugo-server/model"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func start() {
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

	db, err := gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	mysqlDB, err := db.DB()
	utils.HandleErr(err)
	db.AutoMigrate(&model.Admin{}, &model.User{}, &model.Comment{})

	defer mysqlDB.Close()
}
