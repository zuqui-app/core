package server

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal"
	"zuqui-core/internal/domain"
	"zuqui-core/internal/repo"
)

type OnboardRequest struct {
	OTP      string `json:"otp"      xml:"otp"      form:"otp"      validate:"required"`
	Email    string `json:"email"    xml:"email"    form:"email"    validate:"required"`
	Username string `json:"username" xml:"username" form:"username" validate:"required"`
}

type OnboardResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// Onboard new users
func (a *App) AuthOnboard(c *fiber.Ctx) error {
	var req OnboardRequest

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	if err := internal.Validate.Struct(req); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return c.Status(fiber.StatusBadRequest).JSON(internal.ValidationErrorsToMap(errors))
		}
	}

	// Take otp and onboard data
	verified := a.auth.VerifyOTP(req.Email, req.OTP)

	if !verified {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if _, err := a.repo.User.GetUserByEmail(req.Email); !errors.Is(err, repo.ErrNoUser) {
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return fiber.NewError(fiber.StatusBadRequest, "User already exists")
	}

	user, err := a.repo.User.CreateUser(domain.User{
		Email:    req.Email,
		Username: req.Username,
	})
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	access, refresh, err := a.auth.CreateTokenPair(user.Id)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return c.JSON(OnboardResponse{
		Access:  access,
		Refresh: refresh,
	})
}
