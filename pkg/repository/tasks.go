package repository

import (
	"fmt"
	"webtodo/models"
)

var ErrTaskNotFound = fmt.Errorf("task not found")

func AddTask(task *models.Task) (uint, error) {
	if err := GetDBConn().Create(&task).Error; err != nil {
		return 0, err
	}
	return task.Id, nil
}

func GetAllTasks(userId uint) ([]*models.Task, error) {
	var tasks []*models.Task
	if err := GetDBConn().Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func ReassignTask(taskID, userID uint) error {
	if err := GetDBConn().Exec("CALL reassign_task(?, ?)", taskID, userID).Error; err != nil {
		return err
	}
	return nil
}

func GetTaskById(id, userId uint) (*models.Task, error) {
	var task models.Task
	if err := GetDBConn().Where("id = ? AND user_id = ?", id, userId).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func GetExpiredTasksByUser(userId uint) ([]int, error) {
	tasks := []int{}
	if err := GetDBConn().Raw("SELECT get_expired_tasks_by_user(?)", userId).Scan(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// func UpdateTask(id int, userId int) error {
// 	// _, err := findTask(id, userId)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	stmt := "UPDATE tasks SET title = $1, description = $2, completed = $3, user_id = $4, deadline = $5 WHERE id = $6 AND user_id = $7"
// 	_, err = db.GetDBConn().Exec(stmt, task.Title, task.Description, task.IsCompleted, task.UserId, task.Deadline, id, userId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteTask(id, userId uint) error {
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
