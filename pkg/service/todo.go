package service

import (
	"webtodo/models"
	"webtodo/pkg/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo}
}

func (s *TodoService) Add(task *models.Task) (uint, error) {
	return s.repo.Add(task)
}

func (s *TodoService) GetTaskById(id, userId uint) (*models.Task, error) {
	return s.repo.GetTaskById(id, userId)
}

func (s *TodoService) GetAllTasks(userId uint) ([]*models.Task, error) {
	return s.repo.GetAllTasks(userId)
}

func (s *TodoService) UpdateTask(input models.Update, id, userId uint) error {
	return s.repo.UpdateTask(input, id, userId)
}

func (s *TodoService) DeleteTask(id, userId uint) error {
	return s.repo.DeleteTask(id, userId)
}

func (s *TodoService) GetExpiredTasksByUser(userId uint) ([]int, error) {
	return s.repo.GetExpiredTasksByUser(userId)
}

func (s *TodoService) ReassignTask(taskID, userID uint) error {
	return s.repo.ReassignTask(taskID, userID)
}
