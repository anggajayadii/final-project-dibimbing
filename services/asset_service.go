package services

import (
	"asset-management/models"
	"asset-management/repositories"
	"encoding/json"
	"errors"
)

type AssetService interface {
	CreateAsset(userID uint, asset models.Asset) (*models.Asset, error)
	GetAllAssets(userRole models.Role) ([]models.Asset, error)
	GetAssetByID(userRole models.Role, id uint) (*models.Asset, error)
	UpdateAsset(userID uint, userRole models.Role, id uint, updates map[string]interface{}) (*models.Asset, error)
	DeleteAsset(userRole models.Role, id uint) error
}

type assetService struct {
	assetRepo    repositories.AssetRepository
	assetLogRepo repositories.AssetLogRepository
}

func NewAssetService(assetRepo repositories.AssetRepository, assetLogRepo repositories.AssetLogRepository) AssetService {
	return &assetService{
		assetRepo:    assetRepo,
		assetLogRepo: assetLogRepo,
	}
}

func (s *assetService) CreateAsset(userID uint, asset models.Asset) (*models.Asset, error) {
	asset.CreatedBy = userID
	if err := s.assetRepo.Create(&asset); err != nil {
		return nil, err
	}

	// Create initial log
	log := models.AssetLog{
		AssetID:   asset.ID,
		LogType:   models.LogTypeAssignment,
		LogData:   `{"action":"asset_created"}`,
		CreatedBy: userID,
	}
	if err := s.assetLogRepo.Create(&log); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s *assetService) GetAllAssets(userRole models.Role) ([]models.Asset, error) {
	// You can add role-based filtering here if needed
	return s.assetRepo.GetAll()
}

func (s *assetService) GetAssetByID(userRole models.Role, id uint) (*models.Asset, error) {
	// You can add role-based access control here if needed
	return s.assetRepo.GetByID(id)
}

func (s *assetService) UpdateAsset(userID uint, userRole models.Role, id uint, updates map[string]interface{}) (*models.Asset, error) {
	asset, err := s.assetRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Save old values for log
	oldValues := map[string]interface{}{
		"name":        asset.Name,
		"description": asset.Description,
		"status":      asset.Status,
		"location":    asset.Location,
	}

	// Apply updates
	if name, ok := updates["name"].(string); ok {
		asset.Name = name
	}
	if description, ok := updates["description"].(string); ok {
		asset.Description = description
	}
	if status, ok := updates["status"].(string); ok {
		asset.Status = models.AssetStatus(status)
	}
	if location, ok := updates["location"].(string); ok {
		asset.Location = location
	}

	if err := s.assetRepo.Update(asset); err != nil {
		return nil, err
	}

	// Determine log type based on what changed
	var logType models.LogType
	switch {
	case updates["status"] != nil:
		logType = models.LogTypeStatusChange
	case updates["location"] != nil:
		logType = models.LogTypeLocation
	default:
		logType = models.LogTypeAssignment
	}

	// Create log for changes
	logData, _ := json.Marshal(map[string]interface{}{
		"old_values": oldValues,
		"new_values": updates,
	})
	log := models.AssetLog{
		AssetID:   asset.ID,
		LogType:   logType,
		LogData:   string(logData),
		CreatedBy: userID,
	}
	if err := s.assetLogRepo.Create(&log); err != nil {
		return nil, err
	}

	return asset, nil
}

func (s *assetService) DeleteAsset(userRole models.Role, id uint) error {
	if userRole != models.RoleAdmin {
		return errors.New("only admin can delete assets")
	}

	// Optional: Add logging for asset deletion
	// log := models.AssetLog{
	//     AssetID:   id,
	//     LogType:   models.LogTypeAssignment,
	//     LogData:   `{"action":"asset_deleted"}`,
	//     CreatedBy: userID,
	// }
	// if err := s.assetLogRepo.Create(&log); err != nil {
	//     return err
	// }

	return s.assetRepo.Delete(id)
}
