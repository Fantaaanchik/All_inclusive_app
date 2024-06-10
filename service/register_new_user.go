package service

import "all_inclusive_app/models"

func (s *Service) Register(user models.RegisterStruct) error {
	return s.Repository.Register(user)
}