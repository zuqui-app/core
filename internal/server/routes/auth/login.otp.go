package auth

import (
	"bytes"
	"errors"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui/internal/domain"
	"zuqui/internal/repo"
	"zuqui/internal/service/auth"
	"zuqui/internal/service/email"
)

type LoginOTPRequest struct {
	OTP   string `json:"otp"   xml:"otp"   form:"otp"`
	Email string `json:"email" xml:"email" form:"email"`
}

type LoginOTPResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func AuthLoginOTP(ur repo.UserRepo, as auth.Service, es email.Service) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req LoginOTPRequest

		if err := c.BodyParser(&req); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "malformed request")
		}

		if req.Email == "" {
			return fiber.NewError(fiber.StatusBadRequest, "email is required")
		}

		// Send OTP
		if req.OTP == "" {
			otp, err := as.CreateOTP(req.Email)
			if err != nil {
				log.Println(err)
				return fiber.NewError(fiber.StatusInternalServerError)
			}

			user, err := ur.GetUserByEmail(req.Email)
			if err != nil {
				if errors.Is(err, repo.ErrNoUser) {
					sentId, err := SendOTPEmail(es, req.Email, SendOTPProps{
						OTP: otp,
					})
					if err != nil {
						log.Println(err)
						return fiber.NewError(fiber.StatusInternalServerError)
					}
					log.Printf("OTP Email sent: %s\n", sentId)

					return c.SendStatus(fiber.StatusNoContent)
				}

				return fiber.NewError(fiber.StatusInternalServerError)
			}

			sentId, err := SendOTPEmail(es, req.Email, SendOTPProps{
				User: user,
				OTP:  otp,
			})
			if err != nil {
				log.Println(err)
				return fiber.NewError(fiber.StatusInternalServerError)
			}
			log.Printf("OTP Email sent: %s\n", sentId)

			return c.SendStatus(fiber.StatusNoContent)
		}

		// Verify OTP
		verified := as.VerifyOTP(req.Email, req.OTP)

		if !verified {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		user, err := ur.GetUserByEmail(req.Email)
		if err != nil {
			if errors.Is(err, repo.ErrNoUser) {
				return c.Status(fiber.StatusConflict).SendString("onboard required")
			}
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		access, refresh, err := as.CreateTokenPair(user.Id)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		return c.JSON(LoginOTPResponse{
			Access:  access,
			Refresh: refresh,
		})
	}
}

type SendOTPProps struct {
	User domain.User
	OTP  string
}

func SendOTPEmail(
	s email.Service,
	to string,
	props SendOTPProps,
) (string, error) {
	temp, err := template.ParseFiles("internal/service/email/templates/otp_template.html")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := temp.Execute(&buf, props); err != nil {
		return "", err
	}

	html := buf.String()
	req := &email.EmailRequest{
		From:    "Zuqui <zuqui@nichtsam.com>",
		To:      []string{to},
		Subject: "🔑 Here's your OTP for Zuqui!",
		Html:    html,
	}

	sent, err := s.SendEmail(req)
	if err != nil {
		return "", err
	}

	return sent.Id, nil
}
