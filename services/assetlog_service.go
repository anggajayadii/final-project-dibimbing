package services

import (
	"asset-management/models"
	"asset-management/repositories"
	"encoding/json"
)

type AssetLogService interface {
	CreateLog(log *models.AssetLog) error
	GetAssetLogs(assetID uint) ([]models.AssetLog, error)
	GetLogByID(id uint) (*models.AssetLog, error)
	CreateAssetUpdateLog(userID uint, assetID uint, logType models.LogType, oldValues interface{}, newValues interface{}) error
}

type assetLogService struct {
	repo repositories.AssetLogRepository
}

func NewAssetLogService(repo repositories.AssetLogRepository) AssetLogService {
	return &assetLogService{repo: repo}
}

func (s *assetLogService) CreateLog(log *models.AssetLog) error {
	return s.repo.Create(log)
}

func (s *assetLogService) GetAssetLogs(assetID uint) ([]models.AssetLog, error) {
	return s.repo.GetByAssetID(assetID)
}

func (s *assetLogService) GetLogByID(id uint) (*models.AssetLog, error) {
	return s.repo.GetByID(id)
}

func (s *assetLogService) CreateAssetUpdateLog(userID uint, assetID uint, logType models.LogType, oldValues interface{}, newValues interface{}) error {
	logData, err := json.Marshal(map[string]interface{}{
		"old": oldValues,
		"new": newValues,
	})
	if err != nil {
		return err
	}

	log := &models.AssetLog{
		AssetID:   assetID,
		LogType:   logType,
		LogData:   string(logData),
		CreatedBy: userID,
	}

	return s.repo.Create(log)
}
