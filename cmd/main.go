package main

import (
	"go-simple-backend/internal/routes"
	"go-simple-backend/pkg/database"
	"log"
)

func main() {
	// Initialize database connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Gin router
	router := routes.SetupRouter(db)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
