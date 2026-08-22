package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
)

type IUserService interface {
	FindById(userId uint) (*models.User, error)
	Create(request dto.CreateUserInput) (*models.User, error)
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

func (s *UserService) Create(request dto.CreateUserInput) (*models.User, error) {
	newUser := models.User{
		Name: request.Name,
		Email: request.Email,
		Password_digest: request.Password,
		Description: request.Description,
		Icon_url: request.Icon_rul,
	}

	return s.repository.Create(newUser)
}
