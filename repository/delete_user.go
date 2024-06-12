package repository

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) DeleteUser(id int64) error {

	if err := r.db.Table("users").Where("id = ?", id).Delete(models.User{}).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot delete user with id %v, Error: %v", id, err.Error())
		return wrappedErr
	}

	return nil
}
