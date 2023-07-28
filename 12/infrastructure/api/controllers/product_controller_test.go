package controllers

import (
	"bytes"
	"encoding/json"
	application "go-tdd-clean/12/application/product"
	"go-tdd-clean/12/domain/product/mock"
	"go-tdd-clean/12/infrastructure/api/requests"
	"go-tdd-clean/12/infrastructure/api/responses"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/suite"
)

type ProductControllerTest struct {
	suite.Suite
	controller *ProductController
}

func TestProductSuiteStart(t *testing.T) {
	suite.Run(t, new(ProductControllerTest))
}

func (s *ProductControllerTest) SetupTest() {
	repository := mock.NewProductRepositoryFake()
	repository.Seed()
	useCase := application.NewProductUseCase(repository)
	s.controller = NewProductController(useCase)
}

func (s *ProductControllerTest) TestPost_Ok() {
	// Arrange
	app, err := fury.NewWebApplication()
	s.Nil(err)
	app.Post("/api/v1/products", s.controller.Post)

	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	jsonData, err := json.Marshal(request)
	s.Nil(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(jsonData))
	s.Nil(err)

	recorder := httptest.NewRecorder()

	// Act
	app.ServeHTTP(recorder, req)

	// Assert
	s.Equal(http.StatusCreated, recorder.Code)
}

func (s *ProductControllerTest) TestGetMinMaxPrice_Ok() {
	// Arrange
	app, err := fury.NewWebApplication()
	s.Nil(err)
	app.Get("/api/v1/products", s.controller.Get)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/products?min_price=0&max_price=200", nil)
	s.Nil(err)

	recorder := httptest.NewRecorder()

	// Act
	app.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	err = json.NewDecoder(recorder.Body).Decode(&response)
	s.Nil(err)

	s.Equal(http.StatusOK, recorder.Code)
	s.Equal(0, len(response.Errors))
}

func (s *ProductControllerTest) TestGetAll_Ok() {
	// Arrange
	app, err := fury.NewWebApplication()
	s.Nil(err)
	app.Get("/api/v1/products", s.controller.Get)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
	s.Nil(err)

	recorder := httptest.NewRecorder()

	// Act
	app.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	err = json.NewDecoder(recorder.Body).Decode(&response)
	s.Nil(err)

	s.Equal(http.StatusOK, recorder.Code)
	s.Equal(0, len(response.Errors))
}
