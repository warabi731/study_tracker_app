package repositories

import (
	"backend/models"
	"errors"
)

type IUserRepository interface {
	FindById(userId uint) (*models.User, error)
	Create(newUser models.User) (*models.User, error)
}

type UserRepository struct {
	users []models.User
}

func NewUserRepository() IUserRepository {
	dummyUsers := []models.User{
		{ID: 1, Name: "山田 太郎", Email: "yamada@gin.com", Password_digest: "password1", Description: nil, Icon_url: "https://icons/yamada.png"},
		{ID: 2, Name: "佐藤 花子", Email: "sato@gin.com", Password_digest: "password2", Description: nil, Icon_url: "https://icons/sato.png"},
		{ID: 3, Name: "鈴木 一郎", Email: "suzuki@gin.com", Password_digest: "password3", Description: nil, Icon_url: "https://icons/suzuki.png"},
		{ID: 4, Name: "田中 明美", Email: "tanaka@gin.com", Password_digest: "password4", Description: nil, Icon_url: "https://icons/tanaka.png"},
		{ID: 5, Name: "高橋 健", Email: "takahashi@gin.com", Password_digest: "password5", Description: nil, Icon_url: "https://icons/takahashi.png"},
	}

	return &UserRepository{users: dummyUsers}
}

func (r *UserRepository) FindById(userId uint) (*models.User, error) {
	for _, v := range r.users {
		if v.ID == userId {
			return &v, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) Create(newUser models.User) (*models.User, error) {
	newUser.ID = uint(len(r.users) + 1)
	r.users = append(r.users, newUser)

	return &newUser, nil
}