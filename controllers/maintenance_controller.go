package controllers

import (
	"asset-management/constants"
	"asset-management/models"
	"asset-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MaintenanceController struct {
	maintService services.MaintenanceService
}

func NewMaintenanceController(maintService services.MaintenanceService) *MaintenanceController {
	return &MaintenanceController{maintService: maintService}
}

func (ctrl *MaintenanceController) GetRecords(c *gin.Context) {
	records, err := ctrl.maintService.GetAllRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

func (ctrl *MaintenanceController) GetRecordByID(c *gin.Context) {
	id := c.Param("id")
	record, err := ctrl.maintService.GetRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maintenance record not found"})
		return
	}
	c.JSON(http.StatusOK, record)
}

func (ctrl *MaintenanceController) CreateRecord(c *gin.Context) {
	// Dapatkan role user dari context
	userRole, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}

	// Konversi ke type Role
	role := constants.Role(userRole.(models.Role))

	// Cek apakah role diizinkan membuat maintenance record
	allowedRoles := []constants.Role{
		constants.RoleEngineer,
		constants.RoleAdmin,
	}

	roleAllowed := false
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			roleAllowed = true
			break
		}
	}

	if !roleAllowed {
		c.JSON(http.StatusForbidden, gin.H{
			"error":          "your role doesn't have permission to create maintenance records",
			"required_roles": allowedRoles,
		})
		return
	}

	// Lanjutkan proses create maintenance jika diizinkan
	var maint models.Maintenance
	if err := c.ShouldBindJSON(&maint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set user yang membuat record
	userID := c.MustGet("user_id").(uint)
	maint.EngineerID = userID

	if err := ctrl.maintService.CreateRecord(&maint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Maintenance record created successfully",
		"data":    maint,
	})
}
