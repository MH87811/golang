package dto

type CreateProductRequest struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
}

type UpdateProductRequest struct {
	Name  *string `json:"name" binding:"omitempty"`
	Price *uint   `json:"price" binding:"omitempty,gt=0"`
	Stock *uint   `json:"stock" binding:"omitempty,gt=0"`
}
