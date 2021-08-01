package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Uid      uint         `gorm:"primaryKey;autoIncrement" json:"uid"`
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
	IsAdmin  sql.NullBool `gorm:"default:true" json:"isAdmin"`
}
