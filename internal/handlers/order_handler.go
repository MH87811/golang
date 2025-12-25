package handlers

import (
	"shop/internal/models"
	"shop/internal/services"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) Create(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	order, err := h.service.CreateFromCart(user.ID)
	if err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"order": order})
}

func (h *OrderHandler) List(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	orders, err := h.service.ListByUser(user.ID)
	if err != nil {
		Error(c, 500, "error", err.Error())
		return
	}
	Success(c, gin.H{"orders": orders})
}
