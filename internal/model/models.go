package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
}
