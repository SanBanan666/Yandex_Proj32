package models

import "time"

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"event_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRegistration struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"event_id" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // confirmed, cancelled, pending
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
