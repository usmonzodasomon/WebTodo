package repository

import (
	"errors"
	"fmt"
	"webtodo/db"
	"webtodo/models"

	"gorm.io/gorm"
)

var ErrUserNotFound = fmt.Errorf("user not found")

func AddUser(user *models.User) (uint, error) {
	if err := db.GetDBConn().Create(&user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

// func GetAllUsers(userId uint) ([]*models.User, error) {
// 	var Users []*models.User
// 	if err := db.GetDBConn().Where("user_id = ?", userId).Find(&Users).Error; err != nil {
// 		return nil, err
// 	}
// 	return Users, nil
// }

// func GetUserById(id int) (*models.User, error) {
// 	stmt := "SELECT * FROM users WHERE id = $1"
// 	user := models.User{}
// 	err := db.GetDBConn().QueryRow(stmt, id).Scan(&user.Id, &user.Username, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func GetUserTasks(id int) ([]int, error) {
// 	stmt := "SELECT tasks.id FROM tasks JOIN users ON users.id = $1 AND users.id = tasks.user_id"
// 	rows, err := db.GetDBConn().Query(stmt, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var tasks []int
// 	for rows.Next() {
// 		var id int
// 		err := rows.Scan(&id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tasks = append(tasks, id)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return tasks, nil
// }

// func UpdateUser(id int, user *models.User) error {
// 	_, err := findUser(id)
// 	if err != nil {
// 		return err
// 	}
// 	stmt := "UPDATE users SET username = $1, password = $2 WHERE id = $3"
// 	_, err = db.GetDBConn().Exec(stmt, user.Username, user.Password, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func DeleteUser(id int) error {
// 	_, err := findUser(id)
// 	if err != nil {
// 		return err
// 	}
// 	stmt := "DELETE FROM users WHERE id = $1"
// 	_, err = db.GetDBConn().Exec(stmt, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func findUser(id int) (int, error) {
// 	stmt := "SELECT * FROM users WHERE id = $1"
// 	user := &models.User{}
// 	err := db.GetDBConn().QueryRow(stmt, id).Scan(&user.Id, &user.Username, &user.Password)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return 0, ErrTaskNotFound
// 		} else {
// 			return 0, err
// 		}
// 	}
// 	return user.Id, nil
// }

func GetUserByUserAndPassword(username, password string) (uint, error) {
	var user models.User
	if err := db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("incorrect login or password")
		}
		return 0, err
	}
	return user.Id, nil
}
