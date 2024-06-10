package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User представляет модель пользователя
type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name" gorm:"name"`
	Email     string     `json:"email" gorm:"email"`
	Password  string     `json:"password" gorm:"password"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"deleted_at"`
}

// LoginStruct представляет модель логина
type LoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterStruct представляет модель регистрации
type RegisterStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserResponse represents a update user response
type UpdateUserStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims представляет JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
