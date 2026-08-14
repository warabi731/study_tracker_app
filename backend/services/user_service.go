package services

import (
	"backend/models"
	"backend/repositories"
)

type IUserService interface {
	FindById(userId uint) (*models.User, error)
}

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) FindById(userId uint) (*models.User, error) {
	return s.repository.FindById(userId)
}
