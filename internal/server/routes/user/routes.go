package user

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"zuqui/internal/server/middleware"
	"zuqui/internal/service/auth"
)

func RegisterRoutes(
	router fiber.Router,
	as auth.Service,
) {
	userLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	})

	meGroup := router.Group("/me", userLimiter, middleware.AuthMiddleware(as))
	meGroup.Get("/profile", Profile())
	meGroup.Get("/usage", Usage())
}
