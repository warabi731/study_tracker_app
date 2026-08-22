package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
)

type IUserService interface {
	FindById(userId uint) (*models.User, error)
	Create(request dto.CreateUserInput) (*models.User, error)
	Update(userId uint, request dto.UpdateUserInput) (*models.User, error)
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
		Name:            request.Name,
		Email:           request.Email,
		Password_digest: request.Password,
		Description:     request.Description,
		Icon_url:        request.Icon_rul,
	}

	return s.repository.Create(newUser)
}

func (s *UserService) Update(userId uint, request dto.UpdateUserInput) (*models.User, error) {
	target, err := s.FindById(userId)
	if err != nil {
		return nil, err
	}

	if request.Name != nil {
		target.Name = *request.Name
	}
	if request.Email != nil {
		target.Email = *request.Email
	}
	if request.Password != nil {
		target.Password_digest = *request.Password
	}
	if request.Description != nil {
		target.Description = request.Description
	}
	if request.Icon_url != nil {
		target.Icon_url = *request.Icon_url
	}

	return s.repository.Update(*target)
}
