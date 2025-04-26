package controllers

import (
	"asset-management/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssetLogController struct {
	assetLogService services.AssetLogService
}

func NewAssetLogController(assetLogService services.AssetLogService) *AssetLogController {
	return &AssetLogController{assetLogService: assetLogService}
}

// GetAssetLogs godoc
// @Summary Get logs for an asset
// @Security ApiKeyAuth
// @Param asset_id path int true "Asset ID"
// @Success 200 {array} models.AssetLog
// @Router /assets/{asset_id}/logs [get]
func (c *AssetLogController) GetAssetLogs(ctx *gin.Context) {
	assetID, err := strconv.ParseUint(ctx.Param("asset_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset ID"})
		return
	}

	logs, err := c.assetLogService.GetAssetLogs(uint(assetID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

// GetLogByID godoc
// @Summary Get a specific log entry
// @Security ApiKeyAuth
// @Param id path int true "Log ID"
// @Success 200 {object} models.AssetLog
// @Router /logs/{id} [get]
func (c *AssetLogController) GetLogByID(ctx *gin.Context) {
	logID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid log ID"})
		return
	}

	log, err := c.assetLogService.GetLogByID(uint(logID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "log not found"})
		return
	}

	ctx.JSON(http.StatusOK, log)
}
