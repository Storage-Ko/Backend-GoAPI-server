package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Uid      uint         `gorm:"primaryKey;autoIncrement" json:"uid"`
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
	IsAdmin  sql.NullBool `gorm:"default:true" json:"isAdmin"`
}
