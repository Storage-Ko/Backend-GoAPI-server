package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid    uint   `gorm:"primaryKey;autoIncrement" json:"uid"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Grade  int    `json:"grade"`
	Status int    `json:"status"`
}
