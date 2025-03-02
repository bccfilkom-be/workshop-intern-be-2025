package bootstrap

import (
	_productRest "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/interface/rest"
	_productRepo "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/repository"
	_productUsecase "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/usecase"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/jwt"
	_storage "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/infra/storage"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/middleware"

	_userRest "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/interface/rest"
	_userRepo "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/repository"
	_userUsecase "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/usecase"

	_fileRest "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/file/interface/rest"
	_fileUsecase "github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/file/usecase"
)

func MountRoutes(config *Config) {
	jwt := jwt.NewJWT()
	middleware := middleware.NewMiddleware(jwt)
	storage := _storage.New()

	routerGroup := config.Fb.Group("/api/v1")

	productRepository := _productRepo.NewProductMySQL(config.My)
	productUseCase := _productUsecase.NewProductUsecase(productRepository)
	_productRest.NewProductHandler(routerGroup, config.Va, productUseCase, middleware)

	userRepository := _userRepo.NewUserMySQL(config.My)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, jwt)
	_userRest.NewUserHandler(routerGroup, userUseCase)

	fileUseCase := _fileUsecase.NewFileUsecase(storage)
	_fileRest.NewFileHandler(routerGroup, fileUseCase)

}
