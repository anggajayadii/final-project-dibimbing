package repositories

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserByUsername(username string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
