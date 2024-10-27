package database

import (
	"go-simple-backend/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use the POSTGRES_DB connection string
	dsn := os.Getenv("POSTGRES_DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Supabase database:", err)
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}
