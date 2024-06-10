package repo

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot get all users from db, Error description: %s", err.Error())
		return []models.User{}, wrappedErr
	}

	return users, nil
}
