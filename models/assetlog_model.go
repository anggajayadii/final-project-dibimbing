package models

import (
	"time"

	"gorm.io/gorm"
)

type LogType string

const (
	LogTypeStatusChange LogType = "status_change"
	LogTypeAssignment   LogType = "assignment"
	LogTypeLocation     LogType = "location_change"
	LogTypeMaintenance  LogType = "maintenance"
)

type AssetLog struct {
	gorm.Model
	AssetID       uint         `gorm:"not null" json:"asset_id"`
	Asset         Asset        `gorm:"foreignKey:AssetID" json:"asset,omitempty"`
	LogType       LogType      `gorm:"type:enum('status_change','assignment','location_change','maintenance');not null" json:"log_type"`
	LogData       string       `gorm:"type:JSON" json:"log_data"`
	CreatedBy     uint         `gorm:"not null" json:"created_by"`
	User          User         `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
	MaintenanceID *uint        `json:"maintenance_id,omitempty"`
	Maintenance   *Maintenance `gorm:"foreignKey:MaintenanceID" json:"maintenance,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
}
