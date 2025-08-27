package home

import "github.com/gofiber/fiber/v3"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	return c.SendString("Hello")
}
