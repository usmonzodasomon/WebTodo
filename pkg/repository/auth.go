package repository

import (
	"errors"
	"webtodo/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) AddUser(user *models.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (uint, error) {
	var user models.User
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("incorrect login or password")
		}
		return 0, err
	}
	return user.Id, nil
}
