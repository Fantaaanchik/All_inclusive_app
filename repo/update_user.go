package repo

import (
	"all_inclusive_app/models"
	"fmt"
)

func (r *Repository) UpdateUser(id int64, user models.UpdateUserStruct) error {
	if err := r.db.Table("users").Where("id = ?", id).Updates(user).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot update user with id = %v and parameters = %v, Error: %v", id, user, err.Error())
		return wrappedErr
	}
	return nil
}
