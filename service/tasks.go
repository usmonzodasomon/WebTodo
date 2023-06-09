package service

import (
	"webtodo/models"
	"webtodo/pkg/repository"
)

func GetTaskById(id, userId uint) (*models.Task, error) {
	return repository.GetTaskById(id, userId)
}

func GetAllTasks(userId uint) ([]*models.Task, error) {
	return repository.GetAllTasks(userId)
}

func AddTask(task *models.Task) (uint, error) {
	return repository.AddTask(task)
}

func DeleteTask(id, userId uint) error {
	return repository.DeleteTask(id, userId)
}

func GetExpiredTasksByUser(userId uint) ([]int, error) {
	return repository.GetExpiredTasksByUser(userId)
}

func ReassignTask(taskID, userID uint) error {
	return repository.ReassignTask(taskID, userID)
}
