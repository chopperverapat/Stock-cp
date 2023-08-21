package model

import "time"

type User struct {
	ID        uint   `gorm:"primary key"`
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Level     string `gorm:"default:normal"`
	CreatedAt time.Time
}
