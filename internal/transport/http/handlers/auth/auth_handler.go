package auth

import (
	"market/internal/app/users"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *users.App
}

func NewAuthHandler(s *users.App) *AuthHandler {
	return &AuthHandler{service: s}
}

// SignIn
// @Summary Login
// @Tags Auth
// @Description Login
// @ID login-user
// @Accept json
// @Produce json
// @Param input body SignInRequest true "user credentials"
// @Success 200 {object} SignInResponse
// @Failure 400,404,429 {object} object
// @Failure 500 {object} object
// @Failure default {object} object
// @Router /api/auth/sign-in [post]
func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var req SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.service.SignIn(req.Username, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(SignInResponse{
		SignInResponse: *resp,
	})
}
