package handlers

import (
	"shop/internal/dto"
	"shop/internal/models"
	"shop/internal/services"
	"strconv"

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
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		Error(c, 400, "error", "invalid page number")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		Error(c, 400, "error", "invalid limit")
		return
	}

	minPrice, _ := strconv.Atoi(c.DefaultQuery("min_price", "0"))
	maxPrice, _ := strconv.Atoi(c.DefaultQuery("max_price", "0"))
	query := c.Query("q")

	products, total, err := h.service.List(page, limit, minPrice, maxPrice, query)
	if err != nil {
		Error(c, 500, "error", "failed to fetch products")
		return
	}

	Success(c, gin.H{"products": products, "total": total, "limit": limit, "page": page})
}

func (h *ProductHandler) Update(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Error(c, 400, "error", "invalid product id")
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		Error(c, 400, "error", err.Error())
		return
	}

	user, _ := c.Get("user")
	u := user.(models.User)

	updated, err := h.service.Update(uint(productID), u.ID, req)
	if err != nil {
		Error(c, 403, "error", err.Error())
		return
	}

	Success(c, gin.H{"product": updated})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Error(c, 400, "error", "invalid product id")
		return
	}

	user := c.MustGet("user").(models.User)

	if err := h.service.Delete(uint(productID), user.ID); err != nil {
		Error(c, 403, "error", err.Error())
	}

	Success(c, gin.H{"message": "product deleted"})
}
