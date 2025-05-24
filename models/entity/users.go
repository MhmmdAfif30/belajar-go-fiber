package entity

import (
	"time"
)

type User struct {
	Id                   uint      `json:"id" gorm:"primaryKey"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	PasswordConfirmation string    `json:"password_confirmation"`
	CreatedAt            time.Time `json:"created_at" gorm:"autoCreateTime"`

	About            []About      `gorm:"foreignKey:UserID"`
	SentMessages     []Messages `gorm:"foreignKey:SenderID"`
	ReceivedMessages []Messages `gorm:"foreignKey:ReceiverID"`
}
