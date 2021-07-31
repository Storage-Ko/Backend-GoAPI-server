package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	uid       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	id        string
	name      string
	grade     string
	status    int
	createdAt time.Time
}
