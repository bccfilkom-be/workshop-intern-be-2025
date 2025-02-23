package rest

import (
	"net/http"

	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/product/usecase"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Validator      *validator.Validate
	ProductUseCase usecase.ProductUsecaseItf
	middleware     *middleware.Middleware
}

func NewProductHandler(routerGroup fiber.Router, _validator *validator.Validate, productUseCase usecase.ProductUsecaseItf, middleware *middleware.Middleware) {
	productHandler := ProductHandler{
		Validator:      _validator,
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/products")

	routerGroup.Get("/", productHandler.middleware.Authentication, productHandler.GetAllProducts)
	routerGroup.Post("/", productHandler.middleware.Authorization, productHandler.CreateProduct)
	routerGroup.Get("/:id", productHandler.middleware.Authentication, productHandler.GetSpecificProduct)
	routerGroup.Patch("/:id", productHandler.middleware.Authorization,productHandler.UpdateProduct)
	routerGroup.Delete("/:id", productHandler.middleware.Authorization, productHandler.DeleteProduct)
}

func (h *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestCreateProduct
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := h.Validator.Struct(request); err != nil {
		return err
	}

	res, err := h.ProductUseCase.CreateProduct(request)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create product",
		"payload": res,
	})
}

func (h *ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {

	res, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get products",
		"payload": res,
	})
}

func (h *ProductHandler) GetSpecificProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}

	res, err := h.ProductUseCase.GetSpecificProduct(productID)
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get specific product",
		"payload": res,
	})
}

func (h *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestUpdateProduct

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return err
	}

	request.ID = productID

	err = h.ProductUseCase.UpdateProduct(request)
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}

	err = h.ProductUseCase.DeleteProduct(productID)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
