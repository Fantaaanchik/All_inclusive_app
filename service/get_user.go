package service

import (
	"all_inclusive_app/models"
	"fmt"
	"strconv"
)

func (s *Service) GetUser(idS string) (models.User, error) {

	id, err := strconv.Atoi(idS)
	if err != nil {
		wrappedErr := fmt.Errorf("error from strconv on func GetUser service, Error description: %s", err.Error())
		return models.User{}, wrappedErr
	}

	return s.Repository.GetUser(int64(id))
}
