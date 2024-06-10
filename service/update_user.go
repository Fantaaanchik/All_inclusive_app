package service

import (
	"all_inclusive_app/models"
	"fmt"
	"strconv"
)

func (s *Service) UpdateUser(idS string, user models.UpdateUserStruct) error {
	id, err := strconv.Atoi(idS)
	if err != nil {
		wrappedErr := fmt.Errorf("error from strconv on func UpdateUser service, Error description: %s", err.Error())
		return wrappedErr
	}
	return s.Repository.UpdateUser(int64(id), user)
}
