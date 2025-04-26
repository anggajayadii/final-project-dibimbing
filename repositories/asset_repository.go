package repositories

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type AssetRepository interface {
	Create(asset *models.Asset) error
	GetAll() ([]models.Asset, error)
	GetByID(id uint) (*models.Asset, error)
	Update(asset *models.Asset) error
	Delete(id uint) error
	CreateLog(log *models.AssetLog) error
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *assetRepository) GetAll() ([]models.Asset, error) {
	var assets []models.Asset
	err := r.db.Preload("User").Find(&assets).Error
	return assets, err
}

func (r *assetRepository) GetByID(id uint) (*models.Asset, error) {
	var asset models.Asset
	err := r.db.Preload("User").Preload("Maintenances").Preload("Logs").First(&asset, id).Error
	return &asset, err
}

func (r *assetRepository) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *assetRepository) Delete(id uint) error {
	return r.db.Delete(&models.Asset{}, id).Error
}

func (r *assetRepository) CreateLog(log *models.AssetLog) error {
	return r.db.Create(log).Error
}
