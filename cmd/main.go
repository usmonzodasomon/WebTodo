package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"webtodo"
	"webtodo/pkg/handlers"
	"webtodo/pkg/repository"
	"webtodo/pkg/service"

	"webtodo/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	l := logger.GetLogger()
	defer logger.CloseFile()
	if err := initConfig(); err != nil {
		l.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		l.Fatalf("error loading env variables: %s", err.Error())
	}

	repository.StartDbConnection(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	repos := repository.NewRepository(repository.GetDBConn())
	services := service.NewService(repos)
	handler := handlers.NewHandler(services, l)

	MyServer := new(webtodo.Server)
	go func() {
		if err := MyServer.Run(viper.GetString("port"), handler.Routes()); err != nil {
			l.Errorf("Error while starting server %s", err.Error())
			return
		}
	}()
	l.Info("Server is starting...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := repository.CloseDbConnection(); err != nil {
		l.Errorf("error occurred on database connection closing: %s", err.Error())
	}

	l.Info("Shutting down")
	if err := MyServer.Shutdown(context.Background()); err != nil {
		l.Errorf("Error server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
