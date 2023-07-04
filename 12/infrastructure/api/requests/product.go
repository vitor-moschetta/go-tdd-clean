package requests

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type QueryProductFromToDateRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type QueryProductMinMaxPriceRequest struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
