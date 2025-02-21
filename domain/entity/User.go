package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	Email     *string
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	Bids      []Bid  `gorm:"foreignKey:UserID"`
	Roles     []Role `gorm:"many2many:user_roles"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	u.CreatedAt = time.Now().UTC()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now().UTC()
	return
}
