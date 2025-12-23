package dto

type UpdateProductRequest struct {
	Name  *string `json:"name" binding:"omitempty, min=3"`
	Price *uint   `json:"price" binding:"omitempty, gt=0"`
	Stock *uint   `json:"stock" binding:"omitempty, gt=0"`
}
