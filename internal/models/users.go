package models

import (
	auth "movie-app/utils/auth"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Username  string    `gorm:"type:varchar(32);unique"`
	Salt      string    `gorm:"type:longtext""`
	Password  string    `gorm:"type:longtext"`
	Role      string    `gorm:"type:varchar(12)"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime"`
}

func (u *User) ValidPassword(rawtext string) bool {
	return auth.NewVerifyPassword(rawtext, u.Password, u.Salt)
}
