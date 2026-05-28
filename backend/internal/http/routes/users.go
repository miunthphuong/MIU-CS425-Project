package routes

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"miu-cs425-project/backend/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

type createUserRequest struct {
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required,min=8"`
	Role     models.Role `json:"role" binding:"required"`
}

type updateUserRequest struct {
	Name     *string      `json:"name"`
	Email    *string      `json:"email" binding:"omitempty,email"`
	Password *string      `json:"password" binding:"omitempty,min=8"`
	Role     *models.Role `json:"role"`
}

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{db: db}
}

func (handler UserHandler) List(c *gin.Context) {
	var users []models.User
	if err := handler.db.Order("id asc").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (handler UserHandler) Get(c *gin.Context) {
	user, found := handler.findUser(c)
	if !found {
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler UserHandler) Create(c *gin.Context) {
	var request createUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := normalizeRole(request.Role)
	if !role.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role must be supplier, admin, or customer"})
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

	c.JSON(http.StatusCreated, user)
}

func (handler UserHandler) Update(c *gin.Context) {
	user, found := handler.findUser(c)
	if !found {
		return
	}

	var request updateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name != nil {
		user.Name = normalizeName(*request.Name)
	}
	if request.Email != nil {
		user.Email = normalizeEmail(*request.Email)
	}
	if request.Role != nil {
		role := normalizeRole(*request.Role)
		if !role.IsValid() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role must be supplier, admin, or customer"})
			return
		}
		user.Role = role
	}
	if request.Password != nil {
		passwordHash, err := hashPassword(*request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not secure password"})
			return
		}
		user.PasswordHash = passwordHash
	}

	if err := handler.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user could not be updated"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler UserHandler) Delete(c *gin.Context) {
	user, found := handler.findUser(c)
	if !found {
		return
	}

	if err := handler.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user could not be deleted"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (handler UserHandler) findUser(c *gin.Context) (models.User, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return models.User{}, false
	}

	var user models.User
	err = handler.db.First(&user, uint(id)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return models.User{}, false
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load user"})
		return models.User{}, false
	}

	return user, true
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func normalizeRole(role models.Role) models.Role {
	return models.Role(strings.ToLower(strings.TrimSpace(string(role))))
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
