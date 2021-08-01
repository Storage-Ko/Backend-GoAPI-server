package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Pid     uint   `gorm:"primaryKey" json:"pid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Writer  string `json:"writer"`
}
