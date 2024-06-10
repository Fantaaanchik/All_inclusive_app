package repository

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) Login(user models.LoginStruct) (models.User, error) {
	var foundUser models.User
	var err error
	if err = r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&foundUser).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot found user with this parameters, Error description: %s", err.Error())
		return models.User{}, wrappedErr
	}

	return foundUser, nil
}
