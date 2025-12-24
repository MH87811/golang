package handlers

import (
	"shop/internal/models"

	"github.com/gin-gonic/gin"
	"shop/internal/dto"
	"shop/internal/services"
	"strconv"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{service}
}

func (h *CartHandler) Add(c *gin.Context) {
	var req dto.AddCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	user := c.MustGet("user").(models.User)

	if err := h.service.Add(user.ID, req.ProductID, req.Quantity); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"message": "added to cart"})
}

func (h *CartHandler) Update(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Quantity uint `json:"quantity" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	user := c.MustGet("user").(models.User)

	if err := h.service.UpdateItem(user.ID, uint(itemID), req.Quantity); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"message": "cart updated"})
}

func (h *CartHandler) Delete(c *gin.Context) {
	itemID, _ := strconv.Atoi(c.Param("id"))
	user := c.MustGet("user").(models.User)

	if err := h.service.RemoveItem(user.ID, uint(itemID)); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"message": "item removed"})
}

func (h *CartHandler) Get(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	cart, err := h.service.GetCart(user.ID)
	if err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	Success(c, gin.H{"cart": cart})
}
