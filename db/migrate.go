package db

import (
	"github.com/Backend-GoAPI-server/model"
)

func Migrate() {
	db := GetDB()
	db.LogMode(false)

	// Create table when table is not exist
	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
	}
}
