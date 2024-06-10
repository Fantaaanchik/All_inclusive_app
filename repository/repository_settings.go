package repository

import (
	"all_inclusive_app/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	UserRepository models.User
}
