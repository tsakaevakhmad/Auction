package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"uniqueIndex;not null"` // Уникальное название роли
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	return
}
