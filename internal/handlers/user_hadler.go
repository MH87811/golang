package handlers

import (
	"net/http"
	"shop/internal/models"
	"shop/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context, userService *services.UserService) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser, err := userService.Register(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": createdUser,
	})
}
