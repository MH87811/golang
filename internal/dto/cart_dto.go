package dto

type AddCartReq struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required,gt=0"`
}
