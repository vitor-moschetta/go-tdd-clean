package controllers

import (
	"encoding/json"
	productApplication "go-tdd-clean/12/application/product"
	"go-tdd-clean/12/infrastructure/api/requests"
	"go-tdd-clean/12/infrastructure/api/responses"
	"net/http"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type ProductController struct {
	UseCase *productApplication.ProductUseCase
}

func NewProductController(useCase *productApplication.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase: useCase,
	}
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) (err error) {
	var response responses.Response
	var request requests.CreateProductRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return web.NewError(http.StatusBadRequest, "invalid request body")
	}

	input := productApplication.NewCreateProductInput(request.Name, request.Price)
	output := c.UseCase.Create(input)
	response = responses.BuildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) (err error) {
	queryParams := r.URL.Query()

	var input productApplication.GetProductByMinMaxPriceInput

	input.MinPrice, err = strconv.ParseFloat(queryParams.Get("min_price"), 64)
	if err != nil {
		input.MinPrice = 0
	}

	input.MaxPrice, err = strconv.ParseFloat(queryParams.Get("max_price"), 64)
	if err != nil {
		input.MaxPrice = 0
	}

	if input.MinPrice == 0 && input.MaxPrice == 0 {
		output := c.UseCase.GetAll()
		response := responses.BuildResponse(output, r)
		return web.EncodeJSON(w, response, response.StatusCode)
	}

	output := c.UseCase.GetByMinMaxPrice(input)
	response := responses.BuildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}
