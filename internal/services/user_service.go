package services

import (
	"errors"
	"shop/internal/models"
	"shop/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(user models.User) (models.User, error) {
	_, err := s.Repo.FindByEmail(user.Email)
	if err == nil {
		return models.User{}, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	createdUser, err := s.Repo.Save(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}
