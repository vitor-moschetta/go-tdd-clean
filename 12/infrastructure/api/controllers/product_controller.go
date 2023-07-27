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
	response = responses.OutputToResponse(output)

	return web.EncodeJSON(w, response, BuildHttpStatusCode(output, r.Method))
}

func (c *ProductController) GetMinMaxPrice(w http.ResponseWriter, r *http.Request) (err error) {
	var response responses.Response
	var query productApplication.QueryProductMinMaxPrice

	queryParams := r.URL.Query()

	query.MinPrice, err = strconv.ParseFloat(queryParams.Get("min_price"), 64)
	if err != nil {
		query.MinPrice = 0
	}

	query.MaxPrice, err = strconv.ParseFloat(queryParams.Get("max_price"), 64)
	if err != nil {
		query.MaxPrice = 0
	}

	output := c.UseCase.QueryMinMaxPrice2(query)
	response = responses.OutputToResponse(output)

	return web.EncodeJSON(w, response, BuildHttpStatusCode(output, r.Method))
}
