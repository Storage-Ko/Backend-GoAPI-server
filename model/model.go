package model

import (
	"time"
)

type User struct {
	Uid      string    `gorm:"PRIMARY_KEY;Type:varchar(30);UNIQUE;" validate:"len=30"`
	Provider string    `gorm:"Type:varchar(30);default:default"`
	Id       string    `gorm:"Type:varchar(30)"`
	Name     string    `gorm:"Type:varchar(30)"`
	Password string    `gorm:"Type:varchar(255)"`
	Nickname string    `gorm:"Type:varchar(30)"`
	Sex      string    `gorm:"Type:varchar(10)"`
	Birth    time.Time `gorm:"Type:datetime"`
	Phone    int       `gorm:"Type:integer"`
	Date     time.Time `gorm:"Type:datetime"`
}
