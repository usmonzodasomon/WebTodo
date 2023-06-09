package repository

import (
	"webtodo/models"

	"gorm.io/gorm"
)

type Authorization interface {
	AddUser(user *models.User) (uint, error)
	GetUser(username, password string) (uint, error)
}

type Todo interface {
	Add(task *models.Task) (uint, error)
	GetAllTasks(userId uint) ([]*models.Task, error)
	ReassignTask(taskID, userID uint) error
	GetTaskById(id, userId uint) (*models.Task, error)
	GetExpiredTasksByUser(userId uint) ([]int, error)
	UpdateTask(input models.Update, id uint, userId uint) error
	DeleteTask(id, userId uint) error
}

type Repository struct {
	Authorization
	Todo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Todo:          NewTodoPostgres(db),
	}
}
