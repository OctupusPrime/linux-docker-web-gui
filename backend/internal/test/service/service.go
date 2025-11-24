package service

import "linux-docker-web-gui/pkg/models"

type TestRepository interface {
	GetTest() (*models.Test, error)
}

type TestService struct {
	repo TestRepository
}

func NewService(repo TestRepository) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) GetTest() (*models.Test, error) {
	return s.repo.GetTest()
}
