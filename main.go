package main

import (
	"fmt"
	"log"

	config "delivery-validation/config"
	"delivery-validation/pkg/database"
	"delivery-validation/pkg/handlers"
	logger "delivery-validation/pkg/logger"
	"delivery-validation/pkg/router"
)

func main() {
	logger := logger.NewLogger("log.txt")
	logger.InfoLogger.Println("Reading database configuration")

	/*
		databaseConfig, err := config.LoadDatabaseConfiguration()
		if err != nil {
			log.Printf("Error setting database : %s\n", err.Error())
			return
		}
	*/

	//initializing db and router
	logger.InfoLogger.Println("Initializing Program")

	// databaseConfig, err := config.LoadDatabaseConfiguration()
	// if err != nil {
	// 	log.Printf("Error setting database : %s\n", err.Error())
	// 	return

	// }

	database, err := database.NewDatabase("postgres")

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
		return
	}

	router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(database, router, logger)
	handlers.RegisterAllHandlers()
	port := config.LoadPort()
	// port := "8080"
	fmt.Println("Server starting at " + port)
	router.Start(port)
}
