package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Lot struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string `gorm:"type:text; not null;""`
	Description  string `gorm:"type:text; not null"`
	BeginingDate *time.Time
	CreatedDate  time.Time
	Bids         []Bid     `gorm:"foreignKey:LotID"` // Связь с Bid
	Photos       []Photo   `gorm:"foreignKey:LotID"`
	CategoryId   string    `gorm:"type:uuid; not null"`
	Category     *Category `gorm:"foreignKey:CategoryId;references:ID"`
}

func (r *Lot) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == "" {
		r.ID = uuid.NewString()
	}
	r.CreatedDate = time.Now().UTC()
	return
}
