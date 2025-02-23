package bootstrap

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/interface/rest"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/repository"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/usecase"
)

func MountRoutes(config *Config) {
	routerGroup := config.Fb.Group("/api/v1")

	productRepository := repository.NewProductMySQL(config.My)
	productUseCase := usecase.NewProductUsecase(productRepository)
	rest.NewProductHandler(routerGroup, config.Va, productUseCase)
}
