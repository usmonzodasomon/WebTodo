package models

import (
	"time"
)

type Task struct {
	// gorm.Model
	Id          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" binding:"required" gorm:"not null; unique"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	UserId      uint      `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Deadline    uint      `json:"deadline" binding:"required" gorm:"not null"`
	User        User      `json:"-" gorm:"foreignKey:UserId"`
}

// gorm:
