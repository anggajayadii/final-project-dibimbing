package controllers

import (
	"asset-management/middleware"
	"asset-management/models"
	"asset-management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

type RegisterRequest struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required,min=6"`
	FullName string      `json:"full_name" binding:"required"`
	Role     models.Role `json:"role" binding:"required"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Register(req.Username, req.Password, req.FullName, req.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT Token (akan diimplementasikan di middleware)
	token, err := middleware.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
