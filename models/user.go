package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
}

// binding required
