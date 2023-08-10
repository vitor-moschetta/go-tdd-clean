package api

import (
	"encoding/json"
	"go-tdd-clean/50/category/usecase"
	"go-tdd-clean/50/shared"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type CategoryController struct {
	CreateCategoryUseCase usecase.CreateCategoryUseCase
}

func NewProductController(createCategoryUseCase usecase.CreateCategoryUseCase) *CategoryController {
	return &CategoryController{
		CreateCategoryUseCase: createCategoryUseCase,
	}
}

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) (err error) {
	var response shared.Response
	var request CreateCategoryRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return web.NewError(http.StatusBadRequest, "invalid request body")
	}

	input := usecase.CreateCategoryInput{
		Name: request.Name,
	}

	output := c.CreateCategoryUseCase.Execute(input)
	response = shared.BuildResponse(output, r)

	return web.EncodeJSON(w, response, response.StatusCode)
}
