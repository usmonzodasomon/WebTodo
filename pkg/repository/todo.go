package repository

import (
	"fmt"
	"webtodo/models"

	"gorm.io/gorm"
)

var ErrTaskNotFound = fmt.Errorf("task not found")

type TodoPostgres struct {
	db *gorm.DB
}

func NewTodoPostgres(db *gorm.DB) *TodoPostgres {
	return &TodoPostgres{db}
}

func (r *TodoPostgres) Add(task *models.Task) (uint, error) {
	if err := GetDBConn().Create(&task).Error; err != nil {
		return 0, err
	}
	return task.Id, nil
}

func (r *TodoPostgres) GetAllTasks(userId uint) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := GetDBConn().Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TodoPostgres) ReassignTask(taskID, userID uint) error {
	if err := GetDBConn().Exec("CALL reassign_task(?, ?)", taskID, userID).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoPostgres) GetTaskById(id, userId uint) (*models.Task, error) {
	var task models.Task
	if err := GetDBConn().Where("id = ? AND user_id = ?", id, userId).First(&task).Error; err != nil {
		return nil, ErrTaskNotFound
	}
	return &task, nil
}

func (r *TodoPostgres) GetExpiredTasksByUser(userId uint) ([]int, error) {
	tasks := []int{}
	if err := GetDBConn().Raw("SELECT get_expired_tasks_by_user(?)", userId).Scan(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TodoPostgres) UpdateTask(input models.Update, id, userId uint) error {
	task, err := r.GetTaskById(id, userId)
	if err != nil {
		return ErrTaskNotFound
	}
	if input.Title != nil {
		if err := r.db.Model(task).Where("id = ? AND user_id = ?", id, userId).Update("title", input.Title).Error; err != nil {
			return err
		}
	}

	if input.Description != nil {
		if err := r.db.Model(task).Where("id = ? AND user_id = ?", id, userId).Update("description", input.Description).Error; err != nil {
			return err
		}
	}

	if input.IsCompleted != nil {
		if err := r.db.Model(task).Where("id = ? AND user_id = ?", id, userId).Update("is_completed", input.IsCompleted).Error; err != nil {
			return err
		}
	}

	if input.Deadline != nil {
		if err := r.db.Model(task).Where("id = ? AND user_id = ?", id, userId).Update("deadline", input.Deadline).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *TodoPostgres) DeleteTask(id, userId uint) error {
	if err := GetDBConn().Where("user_id = ? AND id = ?", userId, id).Delete(&models.Task{}).Error; err != nil {
		return err
	}
	return nil
}

// func findTask(id, userId uint) (int, error) {
// 	stmt := "SELECT * FROM tasks WHERE id = $1 AND user_id = $2"
// 	task := &models.Task{}
// 	err := db.GetDBConn().QueryRow(stmt, id, userId).Scan(&task.Id, &task.Title, &task.Description, &task.IsCompleted, &task.UserId, &task.CreatedAt, &task.Deadline)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return 0, ErrTaskNotFound
// 		} else {
// 			return 0, err
// 		}
// 	}
// 	return task.Id, nil
// }
