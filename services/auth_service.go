package services

import (
	"asset-management/models"
	"asset-management/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password, fullName string, role models.Role) (*models.User, error)
	Login(username, password string) (*models.User, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(username, password, fullName string, role models.Role) (*models.User, error) {
	// Validasi role
	if role != models.RoleAdmin && role != models.RoleLogistic &&
		role != models.RoleEngineer && role != models.RoleManager {
		return nil, errors.New("role tidak valid")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
		FullName: fullName,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(username, password string) (*models.User, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("username tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("password salah")
	}

	return user, nil
}
