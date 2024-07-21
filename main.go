package main

import (
	"boilerplate/internal/handler"
	"boilerplate/pkg/applog"
	"boilerplate/pkg/database"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("./app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	applog.InitLogger(file)
	applog.Info("Hello world")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	handler.InitHandler(db)
}
