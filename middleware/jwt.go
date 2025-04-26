package middleware

import (
	"asset-management/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) // Ambil dari environment

type Claims struct {
	UserID   uint        `json:"user_id"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
	claims := Claims{
		UserID:           user.ID,
		Username:         user.Username,
		Role:             user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token berlaku 24 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Klaim token tidak valid"})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

func RoleMiddleware(allowedRoles ...models.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Role tidak ditemukan"})
			return
		}

		role, ok := userRole.(models.Role)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Tipe data role tidak valid"})
			return
		}

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Akses ditolak"})
	}
}
