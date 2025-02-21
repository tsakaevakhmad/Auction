package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID       string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string     `gorm:"type:text; not null"`
	ParentID *string    `gorm:"type:uuid"`                         // ParentID может быть NULL (если это корневая категория)
	Parent   *Category  `gorm:"foreignKey:ParentID;references:ID"` // Родительская категория
	Children []Category `gorm:"foreignKey:ParentID"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.NewString()
	}
	return
}
