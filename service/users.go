package service

import (
	"webtodo/models"
	"webtodo/pkg/repository"
)

func AddUser(user *models.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return repository.AddUser(user)
}

// func ChangePassword(password, newPassword string) (error) {

// }

// func GetAllUsers(userId uint) ([]*models.User, error) {
// 	return repository.GetAllUsers(userId)
// }
