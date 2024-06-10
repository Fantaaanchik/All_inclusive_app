package repo

import (
	"log"
)

func (r *Repository) DeleteUser(id int64) error {

	if err := r.db.Delete("user", id).Error; err != nil {
		log.Printf("cannot delete user with id %v, Error: %v", id, err.Error())
		return err
	}

	return nil
}
