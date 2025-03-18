package quiz

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"zuqui/internal/server/middleware"
	"zuqui/internal/service/auth"
	"zuqui/internal/service/quiz"
)

func RegisterRoutes(
	router fiber.Router,
	as auth.Service,
	qs *quiz.Service,
) {
	quizLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Minute,
	})

	quizGroup := router.Group("/quiz", quizLimiter, middleware.AuthMiddleware(as))
	quizGroup.Post("/language/binary/generate", LanguageGenerateBinary(qs))
	quizGroup.Post("/language/scq/generate", LanguageGenerateSCQ(qs))
	quizGroup.Post("/math/binary/generate", MathGenerateBinary())
	quizGroup.Post("/math/scq/generate", MathGenerateSCQ())
}
