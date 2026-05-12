package user

import (
	"market/internal/app/users"

	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	app *users.App
}

func NewUsersHandler(app *users.App) *UsersHandler {
	return &UsersHandler{app}
}

// CreateUser
// @Summary CreateUser User
// @Tags Users
// @Description CreateUser
// @ID create-user
// @Accept json
// @Produce json
// @Param input body createUserReq true "user info"
// @Success 200 {object} createUserResponse
// @Failure 400,404,429 {object} object
// @Failure 500 {object} object
// @Failure default {object} object
// @Router /api/users [post]
func (h *UsersHandler) CreateUser(c *fiber.Ctx) error {
	var req createUserReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	resp, err := h.app.Register(req.Name, req.Email, req.Password, req.PhoneNumbers, req.Role, req.Data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createUserResponse{UUID: resp.ID, Username: resp.Username})
}

// GetCurrentUser
// @Summary Get current user
// @Tags Users
// @Description Get current user
// @ID get-cur-user
// @Accept json
// @Produce json
// @Param input body createUserReq true "user info"
// @Success 200 {object} currentUserResponse
// @Failure 400,404,429 {object} object
// @Failure 500 {object} object
// @Failure default {object} object
// @Router /api/users [get]
func (h *UsersHandler) GetCurrentUser(c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(string)
	if !ok || userId == "" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	resp, err := h.app.GetUserById(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(currentUserResponse{User: *resp})
}
