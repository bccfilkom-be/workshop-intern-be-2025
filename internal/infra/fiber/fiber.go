package fiber

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"time"
)

const idleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout:  idleTimeout,
		ErrorHandler: CustomError,
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("\nGracefully shutting down...")
		_ = app.Shutdown()
	}()

	return app
}

func CustomError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
