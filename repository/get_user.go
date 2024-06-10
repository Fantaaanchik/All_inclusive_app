package repository

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) GetUser(id int64) (models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot get user from db with id = %v, Error description: %s", id, err.Error())
		return models.User{}, wrappedErr
	}

	return user, nil
}
