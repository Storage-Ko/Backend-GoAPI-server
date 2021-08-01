package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Uid       uint      `gorm:"primaryKey;autoIncrement" json:"uid"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	IsAdmin   bool      `sql:"default:true" json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
}
