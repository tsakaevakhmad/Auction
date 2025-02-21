package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string     // A regular string field
	Email     *string    // A pointer to a string, allowing for null values
	Birthday  *time.Time // A pointer to time.Time, can be null
	CreatedAt time.Time  // Automatically managed by GORM for creation time
	UpdatedAt time.Time  // Automatically managed by GORM for update time
	Roles     []Role     `gorm:"many2many:user_roles"`
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
