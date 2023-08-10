package api

import (
	"go-tdd-clean/01/product"
	"go-tdd-clean/12/domain/product/mock"
	"go-tdd-clean/12/infrastructure/api/controllers"
	"go-tdd-clean/50/shared/api"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func Run() error {
	logOptions := []log.Option{
		log.WithCaller(true),
		log.WithStacktraceOnError(true),
		log.WithJSONEncoding(),
	}

	app, err := fury.NewWebApplication([]fury.AppOptFunc{
		fury.WithLogLevel(log.DebugLevel),
		fury.WithLogOptions(logOptions...),
		fury.WithEnableProfiling(),
	}...)

	app.Logger.Info("starting app")
	if err != nil {
		app.Logger.Error("error starting up", log.String("err", err.Error()))
		return err
	}
	app.Router.Use(web.AcceptJSON())
	app.Router.Use(web.Panics())
	app.Router.Use(
		web.LogRequest(
			app.Logger, web.LogRequestConfig{
				IncludeRequest:  true,
				IncludeResponse: true,
			},
		),
	)

	app.Router.ErrorHandler(api.ErrorHandler)

	productRepository := mock.NewProductRepositoryFake()
	productRepository.Seed()
	productUseCase := product.NewProductUseCase(productRepository)
	productController := controllers.NewProductController(productUseCase)

	v1 := app.Router.Group("/api/v1")
	v1.Post("/products", productController.Post)
	v1.Get("/products", productController.Get)
	v1.Get("/error", productController.GetErrorTest)

	return app.Run()
}
