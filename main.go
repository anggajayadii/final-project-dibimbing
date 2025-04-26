package main

import (
	"asset-management/config"
	"asset-management/models"
	"asset-management/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate tables
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Asset{},
		&models.Maintenance{},
		&models.AssetLog{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// Initialize JWT
	config.InitJWTConfig()

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Get port from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
