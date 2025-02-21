package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Photo struct {
	ID    string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Path  string `gorm:"type:text; not null"`
	LotID string `gorm:"type:uuid;not null"`
	Lot   Lot    `gorm:"foreignKey:LotID;references:ID"`
}

func (r *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	return
}
