package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleSupplier Role = "supplier"
	RoleAdmin    Role = "admin"
	RoleCustomer Role = "customer"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Role         Role           `json:"role" gorm:"type:varchar(20);not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (role Role) IsValid() bool {
	switch role {
	case RoleSupplier, RoleAdmin, RoleCustomer:
		return true
	default:
		return false
	}
}
