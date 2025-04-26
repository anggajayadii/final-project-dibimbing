package routes

import (
	"asset-management/config"
	"asset-management/controllers"
	"asset-management/middleware"
	"asset-management/models"
	"asset-management/repositories"
	"asset-management/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Initialize all repositories
	authRepo := repositories.NewAuthRepository(config.DB)
	userRepo := repositories.NewUserRepository(config.DB)
	assetRepo := repositories.NewAssetRepository(config.DB)
	assetLogRepo := repositories.NewAssetLogRepository(config.DB)
	maintenanceRepo := repositories.NewMaintenanceRepository(config.DB)

	// Initialize all services
	authService := services.NewAuthService(authRepo)
	userService := services.NewUserService(userRepo)
	assetService := services.NewAssetService(assetRepo, assetLogRepo)
	assetLogService := services.NewAssetLogService(assetLogRepo)
	maintenanceService := services.NewMaintenanceService(maintenanceRepo)

	// Initialize all controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	assetController := controllers.NewAssetController(assetService)
	assetLogController := controllers.NewAssetLogController(assetLogService)
	maintenanceController := controllers.NewMaintenanceController(maintenanceService)

	// Auth routes (no authentication required)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}

	// Protected routes (require JWT authentication)
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware())
	{
		// User routes
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("", middleware.RoleMiddleware(models.RoleAdmin), userController.GetAllUsers)
			userGroup.GET("/:id", userController.GetUser)
			userGroup.PUT("/:id", userController.UpdateUser)
			userGroup.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), userController.DeleteUser)
		}

		// Asset routes
		assetGroup := apiGroup.Group("/assets")
		{
			assetGroup.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleLogistic), assetController.CreateAsset)
			assetGroup.GET("", assetController.GetAllAssets)
			assetGroup.GET("/:id", assetController.GetAssetByID)
			assetGroup.PUT("/:id", middleware.RoleMiddleware(models.RoleAdmin, models.RoleLogistic), assetController.UpdateAsset)
			assetGroup.DELETE("/:id", middleware.RoleMiddleware(models.RoleAdmin), assetController.DeleteAsset)

			// Asset logs sub-route
			assetGroup.GET("/:id/logs", assetLogController.GetAssetLogs)
		}

		// Maintenance routes
		maintenanceGroup := apiGroup.Group("/maintenances")
		{
			maintenanceGroup.POST("", middleware.RoleMiddleware(models.RoleAdmin, models.RoleEngineer), maintenanceController.CreateRecord)
			maintenanceGroup.GET("", maintenanceController.GetRecords)
			maintenanceGroup.GET("/:id", maintenanceController.GetRecordByID)
		}

		// Log routes
		logGroup := apiGroup.Group("/logs")
		{
			logGroup.GET("/:id", assetLogController.GetLogByID)
		}
	}
}
