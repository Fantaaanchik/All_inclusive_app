package service

import "all_inclusive_app/repository"

type Service struct {
	Repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
