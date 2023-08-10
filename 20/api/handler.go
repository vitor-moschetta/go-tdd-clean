package api

import (
	"encoding/json"
	"go-tdd-clean/20/product"
	"net/http"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type ProductHandler struct {
	CreateProduct           *product.CreateProductUseCase
	GetProductByMinMaxPrice *product.GetProductByMinMaxPriceUseCase
}

func NewProductHandler(
	createProduct *product.CreateProductUseCase,
	getProductByMinMaxPrice *product.GetProductByMinMaxPriceUseCase) *ProductHandler {
	return &ProductHandler{
		CreateProduct:           createProduct,
		GetProductByMinMaxPrice: getProductByMinMaxPrice,
	}
}

func (c *ProductHandler) Post(w http.ResponseWriter, r *http.Request) (err error) {
	var response Response
	var request CreateProductRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return web.NewError(http.StatusBadRequest, "invalid request body")
	}

	input := product.NewCreateProductInput(request.Name, request.Price)
	output := c.CreateProduct.Execute(input)
	response = buildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}

func (c *ProductHandler) Get(w http.ResponseWriter, r *http.Request) (err error) {
	queryParams := r.URL.Query()

	var input product.GetProductByMinMaxPriceInput

	input.MinPrice, err = strconv.ParseFloat(queryParams.Get("min_price"), 64)
	if err != nil {
		input.MinPrice = 0
	}

	input.MaxPrice, err = strconv.ParseFloat(queryParams.Get("max_price"), 64)
	if err != nil {
		input.MaxPrice = 0
	}

	// opcional: fast fail
	if input.MinPrice == 0 && input.MaxPrice == 0 {
		return web.EncodeJSON(w, "invalid query params", http.StatusBadRequest)
	}

	output := c.GetProductByMinMaxPrice.Execute(input)
	response := buildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}
