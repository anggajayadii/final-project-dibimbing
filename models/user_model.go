package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleLogistic Role = "logistik"
	RoleEngineer Role = "engineer"
	RoleManager  Role = "manajer"
)

type User struct {
	gorm.Model
	Username     string        `gorm:"unique;not null" json:"username"`
	Password     string        `gorm:"not null" json:"-"` // `-` untuk skip field di JSON response
	Role         Role          `gorm:"type:enum('admin','logistik','engineer','manajer');not null" json:"role"`
	FullName     string        `json:"full_name,omitempty"` // `omitempty` untuk hide jika kosong
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Assets       []Asset       `gorm:"foreignKey:CreatedBy" json:"assets,omitempty"`
	Maintenances []Maintenance `gorm:"foreignKey:EngineerID" json:"maintenances,omitempty"`
}
