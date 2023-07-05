package controllers

import (
	"encoding/json"
	productApplication "go-tdd-clean/12/application/product"
	"go-tdd-clean/12/infrastructure/api/requests"
	"go-tdd-clean/12/infrastructure/api/responses"
	"net/http"
	"strconv"
)

type ProductController struct {
	UseCase *productApplication.ProductUseCase
}

func NewProductController(useCase *productApplication.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase: useCase,
	}
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	var response responses.Response
	var request requests.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = responses.ItemToResponse(request, "Bad request")
		json.NewEncoder(w).Encode(response)
		return
	}

	input := productApplication.NewCreateProductInput(request.Name, request.Price)
	output := c.UseCase.Create(input)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) GetMinMaxPrice(w http.ResponseWriter, r *http.Request) {
	var response responses.Response
	var query productApplication.QueryProductMinMaxPrice
	var err error

	queryParams := r.URL.Query()

	query.MinPrice, err = strconv.ParseFloat(queryParams.Get("min_price"), 64)
	if err != nil {
		query.MinPrice = 0
	}

	query.MaxPrice, err = strconv.ParseFloat(queryParams.Get("max_price"), 64)
	if err != nil {
		query.MaxPrice = 0
	}

	output := c.UseCase.QueryMinMaxPrice(query)
	response = responses.OutputToResponse(output)
	BuildHttpStatusCode(output, r.Method, w)
	json.NewEncoder(w).Encode(response)
}
