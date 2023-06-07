package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"webtodo"
	"webtodo/db"
	"webtodo/handlers"

	_ "github.com/lib/pq"
)

func main() {
	l := log.New(os.Stdout, "LOG ", log.Ldate|log.Ltime)
	db.StartDbConnection()
	app := handlers.NewTasks(l, db.GetDBConn())
	MyServer := new(webtodo.Server)
	go func() {
		if err := MyServer.Run("9191", app.Routes()); err != nil {
			l.Printf("Error while starting server %s", err.Error())
			return
		}
	}()
	fmt.Println("Server is starting...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	if err := db.CloseDbConnection(); err != nil {
		l.Printf("error occurred on database connection closing: %s", err.Error())
	}

	fmt.Println("Shutting down")
	if err := MyServer.Shutdown(context.Background()); err != nil {
		l.Printf("Error server shutting down: %s", err.Error())
		// comment
	}
}
