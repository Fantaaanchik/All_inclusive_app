package repository

import (
	"fmt"
)

func (r *Repository) DeleteUser(id int64) error {

	if err := r.db.Delete("user", id).Error; err != nil {
		wrappedErr := fmt.Errorf("cannot delete user with id %v, Error: %v", id, err.Error())
		return wrappedErr
	}

	return nil
}
