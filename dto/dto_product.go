package dto

type CreateProdReq struct {
	Name        string  `json:"name" binding:"required, min=3, max=100"`
	Description string  `json:"description" binding:"required, min=3, max=255"`
	Price       float64 `json:"price" binding:"required, gt=0"`
	Stock       int     `json:"stock" binding:"required, gte=0"`
}
