package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bid struct {
	ID        string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Amount    float64 `gorm:"not null"`
	CreatedAt time.Time
	UserID    string `gorm:"type:uuid;not null"` // Внешний ключ на User
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	LotID     string `gorm:"type:uuid;not null"` // Внешний ключ на Lot
	Lot       Lot    `gorm:"foreignKey:LotID;references:ID"`
}

func (b *Bid) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}
	return
}
