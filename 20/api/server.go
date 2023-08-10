package api

import (
	"go-tdd-clean/20/product"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func Run() error {
	app, err := fury.NewWebApplication()

	app.Logger.Info("starting app")
	if err != nil {
		app.Logger.Error("error starting up", log.String("err", err.Error()))
		return err
	}
	app.Router.Use(web.AcceptJSON())
	app.Router.Use(web.Panics())

	productRepo := product.NewProductRepositoryInMemory()
	productRepo.Seed()
	createProductUC := product.NewCreateProductUseCase(productRepo)
	getProductByMinMaxPriceUC := product.NewGetProductByMinMaxPriceUseCase(productRepo)

	handler := NewProductHandler(createProductUC, getProductByMinMaxPriceUC)

	v1 := app.Router.Group("/api/v1")
	v1.Post("/products", handler.Post)
	v1.Get("/products", handler.Get)

	return app.Run()
}
