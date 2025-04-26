package controllers

import (
	"asset-management/dto"
	"asset-management/models"
	"asset-management/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AssetController struct {
	assetService services.AssetService
}

func NewAssetController(assetService services.AssetService) *AssetController {
	return &AssetController{assetService: assetService}
}

// CreateAsset godoc
// @Summary Create a new asset
// @Security ApiKeyAuth
// @Param input body models.Asset true "Asset data"
// @Success 201 {object} models.Asset
// @Router /assets [post]
func (c *AssetController) CreateAsset(ctx *gin.Context) {
	roleVal, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}
	role := roleVal.(models.Role)

	allowedRoles := []models.Role{
		models.RoleAdmin,
		models.RoleLogistic,
	}

	roleAllowed := false
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			roleAllowed = true
			break
		}
	}

	if !roleAllowed {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error":          "your role doesn't have permission to create assets",
			"required_roles": allowedRoles,
		})
		return
	}

	var input dto.CreateAssetInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi serial number tidak boleh kosong
	if input.SerialNumber == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "serial_number is required"})
		return
	}

	purchaseDate, err := time.Parse("2006-01-02", input.PurchaseDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "purchase_date must be in format YYYY-MM-DD"})
		return
	}

	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user_id information not found"})
		return
	}
	userID := userIDVal.(uint)

	asset := models.Asset{
		Name:         input.Name,
		Description:  input.Description,
		PurchaseDate: purchaseDate,
		Location:     input.Location,
		Status:       models.AssetStatus(input.Status),
		CreatedBy:    userID,
		SerialNumber: input.SerialNumber,
	}

	newAsset, err := c.assetService.CreateAsset(userID, asset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Asset created successfully",
		"data":    newAsset,
	})
}

// GetAllAssets godoc
// @Summary Get all assets
// @Security ApiKeyAuth
// @Success 200 {array} models.Asset
// @Router /assets [get]
func (c *AssetController) GetAllAssets(ctx *gin.Context) {
	roleVal, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}
	role := roleVal.(models.Role)

	assets, err := c.assetService.GetAllAssets(role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, assets)
}

// GetAssetByID godoc
// @Summary Get asset by ID
// @Security ApiKeyAuth
// @Param id path int true "Asset ID"
// @Success 200 {object} models.Asset
// @Router /assets/{id} [get]
func (c *AssetController) GetAssetByID(ctx *gin.Context) {
	roleVal, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}
	role := roleVal.(models.Role)

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset ID"})
		return
	}

	asset, err := c.assetService.GetAssetByID(role, uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "asset not found"})
		return
	}

	ctx.JSON(http.StatusOK, asset)
}

// UpdateAsset godoc
// @Summary Update asset
// @Security ApiKeyAuth
// @Param id path int true "Asset ID"
// @Param input body map[string]interface{} true "Update data"
// @Success 200 {object} models.Asset
// @Router /assets/{id} [put]
func (c *AssetController) UpdateAsset(ctx *gin.Context) {
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user_id information not found"})
		return
	}
	userID := userIDVal.(uint)

	roleVal, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}
	role := roleVal.(models.Role)

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset ID"})
		return
	}

	var updates map[string]interface{}
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asset, err := c.assetService.UpdateAsset(userID, role, uint(id), updates)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, asset)
}

// DeleteAsset godoc
// @Summary Delete asset
// @Security ApiKeyAuth
// @Param id path int true "Asset ID"
// @Success 204
// @Router /assets/{id} [delete]
func (c *AssetController) DeleteAsset(ctx *gin.Context) {
	roleVal, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "role information not found"})
		return
	}
	role := roleVal.(models.Role)

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset ID"})
		return
	}

	if err := c.assetService.DeleteAsset(role, uint(id)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
