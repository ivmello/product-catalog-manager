package product

import (
	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	CreateProduct(ctx *fiber.Ctx) error
	ListProducts(ctx *fiber.Ctx) error
}

type handler struct {
	service ProductService
}

func NewProductHandler(service ProductService) ProductHandler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateProduct(ctx *fiber.Ctx) error {
	var input CreateProductInput
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	output, err := h.service.CreateProduct(input)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (h *handler) ListProducts(ctx *fiber.Ctx) error {
	products, err := h.service.ListProducts()
	if err != nil {
		return err
	}
	return ctx.JSON(products)
}
