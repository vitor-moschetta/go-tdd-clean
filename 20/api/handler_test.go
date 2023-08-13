package api

import (
	"bytes"
	"encoding/json"
	"go-tdd-clean/20/product"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/suite"
)

type ProductHandlerTest struct {
	suite.Suite
	ProductHandler *ProductHandler
}

func TestProductControllerSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTest))
}

func (suite *ProductHandlerTest) SetupTest(app *fury.Application, useSeed bool) {
	productRepo := product.NewProductRepositoryInMemory()
	if useSeed {
		productRepo.Seed()
	}
	createProductUseCase := product.NewCreateProductUseCase(productRepo)
	getProductByMinMaxPriceUseCase := product.NewGetProductByMinMaxPriceUseCase(productRepo)
	suite.ProductHandler = NewProductHandler(createProductUseCase, getProductByMinMaxPriceUseCase)
	app.Router.Use(web.AcceptJSON())
	app.Router.Use(web.Panics())

	v1 := app.Router.Group("/api/v1")
	v1.Post("/products", suite.ProductHandler.Post)
	v1.Get("/products", suite.ProductHandler.Get)
}

func (suite *ProductHandlerTest) TestCreateProduct_Ok() {
	// Arrange
	app, err := fury.NewWebApplication()
	suite.NoError(err)

	useSeed := false
	suite.SetupTest(app, useSeed)
	reqBody := CreateProductRequest{
		Name:  "Product 1",
		Price: 100,
	}

	jsonData, err := json.Marshal(reqBody)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(jsonData))
	suite.NoError(err)

	// Act
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, req)

	// Assert
	suite.Equal(http.StatusCreated, recorder.Code)
}

func (suite *ProductHandlerTest) TestCreateProduct_InvalidBody() {
	// Arrange
	app, err := fury.NewWebApplication()
	suite.NoError(err)

	useSeed := false
	suite.SetupTest(app, useSeed)
	reqBody := CreateProductRequest{
		Name:  "",
		Price: -100,
	}

	jsonData, err := json.Marshal(reqBody)
	suite.NoError(err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(jsonData))
	suite.NoError(err)

	// Act
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, req)

	// Assert
	suite.Equal(http.StatusBadRequest, recorder.Code)
}

func (suite *ProductHandlerTest) TestGetProduct_NoContent() {
	// Arrange
	app, err := fury.NewWebApplication()
	suite.NoError(err)

	useSeed := false
	suite.SetupTest(app, useSeed)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products?min_price=100&max_price=200", nil)
	suite.NoError(err)

	// Act
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, req)

	// Assert
	suite.Equal(http.StatusNoContent, recorder.Code)
}

func (suite *ProductHandlerTest) TestGetProduct_Ok() {
	// Arrange
	app, err := fury.NewWebApplication()
	suite.NoError(err)

	useSeed := true
	suite.SetupTest(app, useSeed)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products?min_price=100&max_price=200", nil)
	suite.NoError(err)

	// Act
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, req)

	// Assert
	suite.Equal(http.StatusOK, recorder.Code)
}
