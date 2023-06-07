package db

import (
	"database/sql"
	"log"
	"webtodo/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname = todo_list_db port = 5432 sslmode = disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// CreateTables(db)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Task{})

	return db
}

func StartDbConnection() {
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}

func CloseDbConnection() error {
	db, err := GetDBConn().DB()
	if err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func CreateTables(db *sql.DB) {
	DDLs := []string{
		CreateUsersTable,
		CreateTasksTable,
		CreateGetExpiredTasksByUserFunc,
		CreateReassignTaskProcedure,
	}

	for _, ddl := range DDLs {
		if _, err := db.Exec(ddl); err != nil {
			log.Fatal("Error while creating table. Error is: ", err)
		}
	}
}
