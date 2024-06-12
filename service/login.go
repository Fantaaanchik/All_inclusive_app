package service

import (
	"all_inclusive_app/models"
)

func (s *Service) Login(user models.LoginStruct) (models.User, error) {
	foundedUser, err := s.Repository.Login(user)
	if err != nil {
		return models.User{}, err
	}

	return foundedUser, nil
}
