package repositories

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type AssetLogRepository interface {
	Create(log *models.AssetLog) error
	GetByAssetID(assetID uint) ([]models.AssetLog, error)
	GetByID(id uint) (*models.AssetLog, error)
}

type assetLogRepository struct {
	db *gorm.DB
}

func NewAssetLogRepository(db *gorm.DB) AssetLogRepository {
	return &assetLogRepository{db: db}
}

func (r *assetLogRepository) Create(log *models.AssetLog) error {
	return r.db.Create(log).Error
}

func (r *assetLogRepository) GetByAssetID(assetID uint) ([]models.AssetLog, error) {
	var logs []models.AssetLog
	err := r.db.Preload("User").Preload("Asset").Preload("Maintenance").
		Where("asset_id = ?", assetID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}

func (r *assetLogRepository) GetByID(id uint) (*models.AssetLog, error) {
	var log models.AssetLog
	err := r.db.Preload("User").Preload("Asset").Preload("Maintenance").
		First(&log, id).Error
	return &log, err
}
