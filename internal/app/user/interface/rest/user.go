package rest

import (
	"net/http"

	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/user/usecase"
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase usecase.UserUsecaseItf
}

func NewUserHandler(routerGroup fiber.Router, userUseCase usecase.UserUsecaseItf) {
	UserHandler := UserHandler{
		usecase: userUseCase,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", UserHandler.Register)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		return err
	}

	if err := h.usecase.Register(register); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success register user",
	})
}
