package repository

import (
	"fmt"
	"webtodo/logger"
	"webtodo/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func initDB(cnf *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname = %s port = %s sslmode = %s",
		cnf.Host, cnf.Username, cnf.Password, cnf.DBName, cnf.Port, cnf.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// CreateTables(db)
	if err != nil {
		logger.GetLogger().Error(err)
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Task{})

	return db
}

func StartDbConnection(cnf *Config) {
	database = initDB(cnf)
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
