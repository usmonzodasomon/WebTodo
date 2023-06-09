package service

import (
	"webtodo/models"
	"webtodo/pkg/repository"
)

type Authorization interface {
	AddUser(user *models.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (uint, error)
}

type Todo interface {
	Add(task *models.Task) (uint, error)
	GetAllTasks(userId uint) ([]*models.Task, error)
	ReassignTask(taskID, userID uint) error
	GetTaskById(id, userId uint) (*models.Task, error)
	GetExpiredTasksByUser(userId uint) ([]int, error)
	DeleteTask(id, userId uint) error
}

type Service struct {
	Authorization
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Todo:          NewTodoService(repos.Todo),
	}
}
