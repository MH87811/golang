package services

import (
	"shop/internal/models"
	"shop/internal/repositories"
	"shop/pkg/hash"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user models.User) (models.User, error) {
	_, err := s.repo.FindByEmail(user.Email)
	if err == nil {
		return models.User{}, ErrEmailExists
	}

	hashedPassword, _ := hash.HashPassword(user.Password)
	user.Password = hashedPassword

	return s.repo.Save(user)
}
