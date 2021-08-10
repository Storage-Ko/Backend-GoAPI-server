package model

import (
	"time"
)

type User struct {
	Uid       string    `json:"uid" gorm:"PRIMARY_KEY;Type:varchar(40)"`
	Provider  string    `json:"provider" gorm:"Type:varchar(30);default:'default'"`
	Id        string    `json:"id" gorm:"Type:varchar(30);NOT NULL"`
	Name      string    `json:"name" gorm:"Type:varchar(30);NOT NULL"`
	Password  string    `json:"password" gorm:"Type:varchar(255);NOT NULL"`
	Nickname  string    `json:"nickname" gorm:"Type:varchar(30);NOT NULL"`
	Sex       string    `json:"sex" gorm:"Type:varchar(10);NOT NULL"`
	Birth     time.Time `json:"birth" gorm:"Type:datetime;NOT NULL"`
	Phone     int       `json:"phone" gorm:"Type:integer;NOT NULL"`
	CreatedAt time.Time `json:"createdAt" gorm:"Type:datetime;NOT NULL"`
}

func (User) TableName() string {
	return "user"
}
