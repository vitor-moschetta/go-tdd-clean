package api

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
