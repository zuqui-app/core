package server

import (
	"bytes"
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/service/email"
)

type LoginOTPInput struct {
	OTP   string `json:"otp"   xml:"otp"   form:"otp"`
	Email string `json:"email" xml:"email" form:"email"`
}

// Login with email (or sms) OTP
func (a *App) AuthLoginOTP(c *fiber.Ctx) error {
	// Take email address or otp
	input := new(LoginOTPInput)

	if err := c.BodyParser(input); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "malformed otp login input")
	}

	if input.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	// Send OTP
	if input.OTP == "" {
		otp, err := a.auth.CreateOTP(input.Email)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		sentId, err := SendOTPEmail(a.email, input.Email, SendOTPProps{
			Username: "", // get user if exist
			OTP:      otp,
		})
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}
		log.Printf("OTP Email sent: %s\n", sentId)

		return c.SendStatus(fiber.StatusNoContent)
	}

	// Verify OTP
	verified := a.auth.VerifyOTP(input.Email, input.OTP)

	if !verified {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	username := ""

	if username == "" {
		return c.Status(fiber.StatusConflict).SendString("onboard required")
	}

	return c.SendString("tokens")
}

type SendOTPProps struct {
	Username string
	OTP      string
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
