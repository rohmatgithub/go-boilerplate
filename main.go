package main

import (
	"boilerplate/internal/handler"
	"boilerplate/pkg/applog"
	"boilerplate/pkg/configs"
	"boilerplate/pkg/database"
	"boilerplate/pkg/util"
	"log"
)

func main() {

	err := configs.Init()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}

	f, err := applog.InitLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer f.Close()

	db, err := database.ConnectAndMigratePostgres()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	applog.Info().Msg("Application started")
	util.InitializeI18n()
	handler.InitHandler(db)
}
