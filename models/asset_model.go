package models

import (
	"time"

	"gorm.io/gorm"
)

type AssetStatus string

const (
	AssetOK               AssetStatus = "OK"
	AssetUnderMaintenance AssetStatus = "Under Maintenance"
	AssetDismantle        AssetStatus = "Dismantle"
)

type CreateAssetInput struct {
	Name               string `json:"name" binding:"required"`
	PurchaseDate       string `json:"purchase_date" binding:"required"`
	WarrantyExpiryDate string `json:"warranty_expiry_date"`
}

type Asset struct {
	gorm.Model
	Name         string        `gorm:"not null" json:"name"`
	Description  string        `json:"description,omitempty"`
	SerialNumber string        `gorm:"unique" json:"serial_number"`
	PurchaseDate time.Time     `json:"purchase_date,omitempty"`
	Status       AssetStatus   `gorm:"type:enum('OK','Under Maintenance','Dismantle');default:'OK'" json:"status"`
	Location     string        `json:"location,omitempty"`
	CreatedBy    uint          `gorm:"not null" json:"created_by"`
	User         User          `gorm:"foreignKey:CreatedBy" json:"user,omitempty"`
	Maintenances []Maintenance `gorm:"foreignKey:AssetID" json:"maintenances,omitempty"`
	Logs         []AssetLog    `gorm:"foreignKey:AssetID" json:"logs,omitempty"`
}
