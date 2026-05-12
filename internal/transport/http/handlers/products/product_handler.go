package products

import (
	"market/internal/app/products"

	"github.com/gofiber/fiber/v2"
)

type ProductsHandler struct {
	app *products.App
}

func NewProductsHandler(s *products.App) *ProductsHandler {
	return &ProductsHandler{app: s}
}

// CreateProduct
// @Summary CreateProduct
// @Tags Products
// @Description CreateProduct
// @ID create-product
// @Accept json
// @Produce json
// @Param input body createUserReq true "Product info"
// @Success 200 {object} createUserResponse
// @Failure 400,404,429 {object} object
// @Failure 500 {object} object
// @Failure default {object} object
// @Router /api/products [post]
func (h *ProductsHandler) CreateProduct(c *fiber.Ctx) error {
	var req createProductReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := h.app.Register(req.Name, req.MeasurementId, req.ParentId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}
