package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	pid       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	title     string
	content   string
	writer    string
	createdAt time.Time
}
