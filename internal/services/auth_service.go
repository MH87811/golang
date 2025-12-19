package services

import (
	"time"

	"shop/internal/repositories"
	"shop/pkg/hash"
	"shop/pkg/jwtpkg"
)

type AuthService struct {
	repo repositories.UserRepository
	jwt  *jwtpkg.JWT
}

func NewAuthService(repo repositories.UserRepository, jwt *jwtpkg.JWT) *AuthService {
	return &AuthService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", ErrInvalidCredential
	}

	if !hash.CheckPassword(password, user.Password) {
		return "", "", ErrInvalidCredential
	}

	access, _ := s.jwt.Generate(user, "access", time.Minute*15)
	refresh, _ := s.jwt.Generate(user, "refresh", time.Hour*24*7)

	return access, refresh, nil
}

func (s *AuthService) Refresh(refreshToken string) (string, error) {
	claims, err := s.jwt.VerifyRefresh(refreshToken)
	if err != nil {
		return "", err
	}

	user, err := s.repo.FindByID(claims.UserID)
	if err != nil {
		return "", err
	}

	return s.jwt.Generate(user, "access", time.Minute*15)
}
