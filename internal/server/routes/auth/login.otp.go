package auth

import "github.com/gofiber/fiber/v2"

// Login with email (or sms) OTP
func LoginOTP(c *fiber.Ctx) error {
	// Take email address or otp
	// if otp attached
	//   get email out of top
	//   if email user exist
	//     return tokens
	//   else
	//     return onboard required
	// else if email address attached
	//   send otp email
	// else
	//   error

	return c.SendStatus(fiber.StatusNotImplemented)
}
