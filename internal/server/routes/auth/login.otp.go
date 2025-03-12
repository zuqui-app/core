package auth

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v2"

	"zuqui-core/internal/email"
	"zuqui-core/internal/services"
)

type LoginOTPInput struct {
	OTP   string `json:"otp"   xml:"otp"   form:"otp"`
	Email string `json:"email" xml:"email" form:"email"`
}

// Login with email (or sms) OTP
func LoginOTP(c *fiber.Ctx) error {
	// Take email address or otp
	input := new(LoginOTPInput)

	if err := c.BodyParser(input); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "malformed otp login input")
	}

	if input.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	// Send otp
	if input.OTP == "" {
		otp := createOTP()

		if err := services.RedisClient.Set(context.Background(), input.Email, otp, 10*time.Minute).Err(); err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}

		sent, err := email.SendOTPEmail(input.Email, email.SendOTPProps{
			Username: "", // get user if exist
			OTP:      otp,
		})
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError)
		}
		log.Printf("OTP Email sent: %s\n", sent.Id)

		return c.SendStatus(fiber.StatusNoContent)
	}

	// Verify otp
	data := services.RedisClient.Get(context.Background(), input.Email)

	if input.OTP != data.Val() {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	username := ""

	if username == "" {
		return c.Status(fiber.StatusConflict).SendString("onboard required")
	}

	return c.SendString("tokens")
}

func createOTP() string {
	chars := "ABCDEFGHJKLMNPQRTUVWXY346789" // Without O, 0, I, 1, S, 5, Z, 2
	r := make([]byte, 6)
	for i := range r {
		r[i] = chars[rand.IntN(len(chars))]
	}

	return string(r)
}
