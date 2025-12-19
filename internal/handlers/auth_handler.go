package handlers

import (
	"errors"
	"shop/internal/dto"
	"shop/internal/models"
	"shop/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
	authService *services.AuthService
}

func NewAuthHandler(u *services.UserService, a *services.AuthService) *AuthHandler {
	return &AuthHandler{u, a}
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		Error(c, 401, "NO_REFRESH_TOKEN", "refresh token invalid")
	}

	access, err := h.authService.Refresh(refreshToken)
	if err != nil {
		Error(c, 401, "INVALID_REFRESH", err.Error())
		return
	}

	Success(c, gin.H{"access_token": access})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		Error(c, 400, "VALIDATION_ERROR", err.Error())
		return
	}

	user, err := h.userService.Register(models.User{Email: req.Email, Password: req.Password})
	if errors.Is(err, services.ErrEmailExists) {
		Error(c, 409, "EMAIL_EXISTS", err.Error())
		return
	}

	Success(c, gin.H{"id": user.ID, "email": user.Email})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		Error(c, 400, "VALIDATION_ERROR", err.Error())
		return
	}

	access, refresh, err := h.authService.Login(req.Email, req.Password)
	if errors.Is(err, services.ErrInvalidCredential) {
		Error(c, 401, "INVALID_CREDENTIALS", err.Error())
		return
	}

	c.SetCookie("refresh_token", refresh, 7*24*60*60, "/", "", false, true)
	Success(c, gin.H{"access_token": access, "refresh_token": refresh})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	user, _ := c.Get("user")
	Success(c, gin.H{"user": user})
}
