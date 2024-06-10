package service

import "all_inclusive_app/repo"

type Service struct {
	Repository *repo.Repository
}

func NewService(repo *repo.Repository) *Service {
	return &Service{Repository: repo}
}
