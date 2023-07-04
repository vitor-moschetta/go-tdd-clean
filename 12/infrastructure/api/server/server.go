package server

import (
	"log"
	"net/http"

	"go-tdd-clean/12/application/product"
	"go-tdd-clean/12/domain/product/mock"

	"go-tdd-clean/12/infrastructure/api/controllers"

	"github.com/gorilla/mux"
	"github.com/vitormoschetta/go/pkg/middlewares"
)

func Start() {

	productRepository := mock.NewProductRepositoryFake()
	productRepository.Seed()
	productUseCase := product.NewProductUseCase(productRepository)
	productController := controllers.NewProductController(productUseCase)

	router := mux.NewRouter()
	router.Use(middlewares.AcceptJSON)

	router.HandleFunc("/api/v1/products", productController.Post).Methods(http.MethodPost)

	router.HandleFunc("/api/v1/products", productController.GetFromToDate).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/products", productController.GetMinMaxPrice).Methods(http.MethodGet)

	port := "8080"
	log.Println("Listening on port", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
