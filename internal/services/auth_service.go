package services

import (
	"errors"
	"time"

	"shop/internal/auth"
	"shop/internal/config"
	"shop/internal/models"
	"shop/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repositories.UserRepository
	Cfg  *config.Config
}

func NewAuthService(repo repositories.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		Repo: repo,
		Cfg:  cfg,
	}
}

func (s *AuthService) GenerateToken(user models.User, tokenType string, ttl int64) (string, error) {
	claims := auth.TokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Cfg.Issuer,
			Subject:   "user-auth",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Cfg.JWTSecret))
}

func (s *AuthService) Refresh(refreshToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&auth.TokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(s.Cfg.JWTSecret), nil
		},
	)

	claims, ok := token.Claims.(*auth.TokenClaims)
	if err != nil || !ok || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	if claims.Type != "refresh" {
		return "", "", errors.New("invalid token type")
	}

	user, err := s.Repo.FindByID(claims.UserID)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	newAccess, err := s.GenerateToken(user, "access", s.Cfg.AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	newRefresh, err := s.GenerateToken(user, "refresh", s.Cfg.RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return newAccess, newRefresh, nil
}

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := s.GenerateToken(user, "access", s.Cfg.AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.GenerateToken(user, "refresh", s.Cfg.RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
