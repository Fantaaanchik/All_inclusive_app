package repo

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) Register(user models.RegisterStruct) error {
	if err := r.db.Table("users").Create(&user).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot create user, Error description: %s", err.Error())
		return wrappedErr
	}
	return nil
}
