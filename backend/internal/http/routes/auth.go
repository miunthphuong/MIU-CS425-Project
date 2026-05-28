package routes

import (
	"errors"
	"net/http"

	"miu-cs425-project/backend/internal/auth"
	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db         *gorm.DB
	jwtService auth.JWTService
}

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type authResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func NewAuthHandler(db *gorm.DB, jwtService auth.JWTService) AuthHandler {
	return AuthHandler{
		db:         db,
		jwtService: jwtService,
	}
}

func (handler AuthHandler) RegisterForRole(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request registerRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		passwordHash, err := hashPassword(request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not secure password"})
			return
		}

		user := models.User{
			Name:         normalizeName(request.Name),
			Email:        normalizeEmail(request.Email),
			PasswordHash: passwordHash,
			Role:         role,
		}

		if err := handler.db.Create(&user).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists or could not be created"})
			return
		}

		token, err := handler.jwtService.Generate(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
			return
		}

		c.JSON(http.StatusCreated, authResponse{Token: token, User: user})
	}
}

func (handler AuthHandler) LoginForRole(role models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request loginRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		email := normalizeEmail(request.Email)
		err := handler.db.Where("email = ? AND role = ?", email, role).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load user"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}

		token, err := handler.jwtService.Generate(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
			return
		}

		c.JSON(http.StatusOK, authResponse{Token: token, User: user})
	}
}
