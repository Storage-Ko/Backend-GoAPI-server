package main

import (
	"fmt"

	"github.com/Backend-GoAtreugo-server/model"
	"github.com/Backend-GoAtreugo-server/utils"
	"github.com/joho/godotenv"
	"github.com/savsgio/atreugo/v11"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	var dbConfig map[string]string
	dbConfig, err := godotenv.Read()
	utils.HandleErr(err)

	config := atreugo.Config{
		Addr: "0.0.0.0:4030",
	}
	server := atreugo.New(config)
	v1 := server.NewGroupPath("/v1")

	v1.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello world")
	})

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
	/*
		admin := model.Admin{
			Id:        "Admin",
			Name:      "root",
			Password:  "root",
			CreatedAt: time.Now(),
		}

	*/
	//db.Select("id", "name", "password").Create(&admin)

	err = server.ListenAndServe()
	utils.HandleErr(err)

	defer mysqlDB.Close()
}
