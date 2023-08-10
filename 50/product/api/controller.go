package api

import (
	"encoding/json"

	"go-tdd-clean/50/product/usecase"
	"go-tdd-clean/50/shared"
	"go-tdd-clean/50/shared/repository"
	"net/http"
	"strconv"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type ProductController struct {
	CreateProductUseCase           usecase.CreateProductUseCase
	GetProductByMinMaxPriceUseCase usecase.GetProductByMinMaxPriceUseCase
	repository.RepositoryContainer
}

func NewProductController(
	createProductUseCase usecase.CreateProductUseCase,
	getProductByMinMaxPriceUseCase usecase.GetProductByMinMaxPriceUseCase,
	repositoryContainer repository.RepositoryContainer,
) *ProductController {
	return &ProductController{
		CreateProductUseCase:           createProductUseCase,
		GetProductByMinMaxPriceUseCase: getProductByMinMaxPriceUseCase,
		RepositoryContainer:            repositoryContainer,
	}
}

func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) (err error) {
	var response shared.Response
	var request CreateProductRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return web.NewError(http.StatusBadRequest, "invalid request body")
	}

	input := usecase.CreateProductInput{
		Name:  request.Name,
		Price: request.Price,
	}

	output := c.CreateProductUseCase.Execute(input)
	response = shared.BuildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}

func (c *ProductController) Get(w http.ResponseWriter, r *http.Request) (err error) {
	queryParams := r.URL.Query()

	var input usecase.GetProductByMinMaxPriceInput

	input.MinPrice, err = strconv.ParseFloat(queryParams.Get("min_price"), 64)
	if err != nil {
		input.MinPrice = 0
	}

	input.MaxPrice, err = strconv.ParseFloat(queryParams.Get("max_price"), 64)
	if err != nil {
		input.MaxPrice = 0
	}

	output := c.GetProductByMinMaxPriceUseCase.Execute(input)
	response := shared.BuildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}
