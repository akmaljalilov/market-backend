package http

import (
	"market/internal/app/users"

	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	service *users.Service
}

func NewUsersHandler(s *users.Service) *UsersHandler {
	return &UsersHandler{service: s}
}

type createUserResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}
type createUserReq struct {
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	PhoneNumbers []string   `json:"phoneNumbers"`
	Role         users.Role `json:"role"`
	Data         string     `json:"data"`
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

	resp, err := h.service.Register(req.Name, req.Email, req.Password, req.PhoneNumbers, req.Role, req.Data)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createUserResponse{UUID: resp.ID, Username: resp.Username})
}
