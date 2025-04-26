package main

import (
	"asset-management/config"
	"asset-management/models"
	"asset-management/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Menghubungkan ke database
	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Asset{},
		&models.Maintenance{},
		&models.AssetLog{},
	)

	// Auto migrate di main.go
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// Initialize JWT configuration
	config.InitJWTConfig()

	router := gin.Default()
	routes.SetupRoutes(router)

	// Menjalankan server di port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
