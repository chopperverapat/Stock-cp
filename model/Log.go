package model

import "time"

type Log struct {
	ID         uint      `gorm:"primaryKey"`
	Username   string    `gorm:"not null"`
	ClientIP   string    `gorm:"not null"`
	RequestURI string    `gorm:"not null"`
	Method     string    `gorm:"not null"`
	Path       string    `gorm:"not null"`
	StatusCode int       `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null"`
}
