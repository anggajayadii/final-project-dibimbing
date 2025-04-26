package models

import (
	"gorm.io/gorm"
)

type MaintenanceStatus string

const (
	MaintenanceComplete            MaintenanceStatus = "Complete"
	MaintenanceCancelled           MaintenanceStatus = "Cancelled"
	MaintenanceNeedPartReplacement MaintenanceStatus = "Need Part Replacement"
)

type MaintenanceType string

const (
	MaintenancePreventive   MaintenanceType = "Preventive Maintenance"
	MaintenanceTroubleshoot MaintenanceType = "Troubleshoot"
)

type Maintenance struct {
	gorm.Model
	AssetID         uint              `gorm:"not null" json:"asset_id"`
	Asset           Asset             `gorm:"foreignKey:AssetID" json:"asset,omitempty"`
	EngineerID      uint              `gorm:"not null" json:"engineer_id"`
	Engineer        User              `gorm:"foreignKey:EngineerID" json:"engineer,omitempty"`
	MaintenanceDate DateOnly          `gorm:"not null" json:"maintenance_date"`
	CompletionDate  DateOnly          `json:"completion_date,omitempty"`
	Description     string            `gorm:"not null" json:"description"`
	Status          MaintenanceStatus `gorm:"type:enum('Complete','Cancelled','Need Part Replacement');not null" json:"status"`
	Type            MaintenanceType   `gorm:"type:enum('Preventive Maintenance','Troubleshoot');not null" json:"type"`
	PartsNeeded     string            `gorm:"type:text" json:"parts_needed,omitempty"`
	Notes           string            `json:"notes,omitempty"`
}
