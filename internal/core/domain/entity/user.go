package entity

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	Birthday    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Bids        []Bid                 `gorm:"foreignKey:UserID"`
	Roles       []Role                `gorm:"many2many:user_roles"`
	Credentials []webauthn.Credential `gorm:"type:jsonb;serializer:json"`
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

func (u User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u User) WebAuthnName() string {
	return u.Email
}

func (u User) WebAuthnDisplayName() string {
	return u.Name
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u User) WebAuthnIcon() string {
	return "" // Можно оставить пустым или добавить ссылку на аватар
}
