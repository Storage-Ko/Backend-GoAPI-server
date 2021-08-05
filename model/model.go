package model

import (
	"time"
)

type User struct {
	Uid      string    `gorm:"PRIMARY_KEY;Type:varchar(40)"`
	Provider string    `gorm:"Type:varchar(30);default:'default'"`
	Id       string    `gorm:"Type:varchar(30);NOT NULL"`
	Name     string    `gorm:"Type:varchar(30);NOT NULL"`
	Password string    `gorm:"Type:varchar(255);NOT NULL"`
	Nickname string    `gorm:"Type:varchar(30);NOT NULL"`
	Sex      string    `gorm:"Type:varchar(10);NOT NULL"`
	Birth    time.Time `gorm:"Type:datetime;NOT NULL"`
	Phone    int       `gorm:"Type:integer;NOT NULL"`
	Date     time.Time `gorm:"Type:datetime;NOT NULL"`
}
