package auth

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"zuqui/internal"
	"zuqui/internal/domain"
	"zuqui/internal/repo"
	"zuqui/internal/service/auth"
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
func AuthOnboard(ur repo.UserRepo, as auth.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
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
		verified := as.VerifyOTP(req.Email, req.OTP)

		if !verified {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if _, err := ur.GetUserByEmail(req.Email); !errors.Is(err, repo.ErrNoUser) {
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError)
			}

			return fiber.NewError(fiber.StatusBadRequest, "User already exists")
		}

		user, err := ur.CreateUser(domain.User{
			Email:    req.Email,
			Username: req.Username,
		})
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		access, refresh, err := as.CreateTokenPair(user.Id)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.JSON(OnboardResponse{
			Access:  access,
			Refresh: refresh,
		})
	}
}
