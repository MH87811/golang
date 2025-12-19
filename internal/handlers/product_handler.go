package handlers

import (
	"shop/internal/models"
	"shop/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	userRaw, exists := c.Get("user")
	if !exists {
		Error(c, 401, "unauthorized", "user not found in context")
		return
	}

	user, ok := userRaw.(models.User)
	if !ok {
		Error(c, 500, "error", "invalid user type in context")
		return
	}

	product.UserID = user.ID

	created, err := h.service.Create(product)
	if err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"product": created})
}

func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.service.List()
	if err != nil {
		Error(c, 500, "internal error", err.Error())
		return
	}

	Success(c, gin.H{"products": products})
}
