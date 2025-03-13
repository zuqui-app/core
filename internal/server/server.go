package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/generative-ai-go/genai"
	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v2"

	"zuqui-core/internal/email"
	"zuqui-core/internal/quiz"
)

type FiberServer struct {
	*fiber.App
	email *email.EmailService
	quiz  *quiz.Generator
	redis *redis.Client
}

func New(resend *resend.Client, genai *genai.Client, redis *redis.Client) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "zuqui-core",
			AppName:      "zuqui-core",
		}),

		email: email.New(resend),
		quiz:  quiz.New(genai),
		redis: redis,
	}

	return server
}

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Use(logger.New())
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	quizGroup := s.Group("/quiz")
	quizGroup.Post("/language/binary/generate", s.LanguageGenerateBinary)
	quizGroup.Post("/language/scq/generate", s.LanguageGenerateSCQ)
	quizGroup.Post("/math/binary/generate", s.MathGenerateBinary)
	quizGroup.Post("/math/scq/generate", s.MathGenerateSCQ)

	authGroup := s.Group("/auth")
	authGroup.Post("/login/otp", s.AuthLoginOTP)
	authGroup.Post("/webauthn/registration", s.WebAuthnRegistration)
	authGroup.Post("/webauthn/authentication", s.WebAuthnAuthentication)

	meGroup := s.Group("/me")
	meGroup.Get("/profile", s.MeProfile)
	meGroup.Get("/usage", s.MeUsage)
}
