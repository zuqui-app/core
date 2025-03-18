package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"zuqui/internal/repo"
	"zuqui/internal/service/auth"
	"zuqui/internal/service/email"
	"zuqui/internal/service/quiz"
)

func RegisterRoutes(
	router fiber.Router,
	ur repo.UserRepo,
	as auth.Service,
	es email.Service,
	qs *quiz.Service,
) {
	authLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 2 * time.Minute,
	})

	authGroup := router.Group("/auth", authLimiter)
	authGroup.Post("/login/otp", AuthLoginOTP(ur, as, es))
	authGroup.Post("/onboard", AuthOnboard(ur, as))
	authGroup.Post("/token/refresh", AuthTokenRefresh(as))
	authGroup.Post("/webauthn/registration", WebAuthnRegistration())
	authGroup.Post("/webauthn/authentication", WebAuthnAuthentication())
}
