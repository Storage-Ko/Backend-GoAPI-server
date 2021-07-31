package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	uid       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	id        string
	name      string
	password  string
	isAdmin   bool `gorm:"default:true"`
	createdAt time.Time
}
