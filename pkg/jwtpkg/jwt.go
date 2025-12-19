package jwtpkg

import (
	"errors"
	"shop/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret string
}

type Claims struct {
	UserID uint
	Type   string
}

func New(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Generate(user models.User, tokenType string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"type":    tokenType,
		"exp":     time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

func (j *JWT) Verify(tokenStr string) (*Claims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	mapClaims := token.Claims.(jwt.MapClaims)

	return &Claims{
		UserID: uint(mapClaims["user_id"].(float64)),
		Type:   mapClaims["type"].(string),
	}, nil
}

func (j *JWT) VerifyRefresh(token string) (*Claims, error) {
	claims, err := j.Verify(token)
	if err != nil {
		return nil, err
	}

	if claims.Type != "refresh" {
		return nil, errors.New("not refresh token")
	}

	return claims, nil
}
