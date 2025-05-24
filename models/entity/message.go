package entity

import (
	"time"
)

type Messages struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SenderID   uint      `json:"sender_id"`
	ReceiverID uint      `json:"receiver_id"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`

	Sender   User `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Receiver User `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
