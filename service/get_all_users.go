package service

import "all_inclusive_app/models"

func (s *Service) GetAllUsers() ([]models.User, error) {
	return s.Repository.GetAllUsers()
}
