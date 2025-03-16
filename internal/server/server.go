package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"zuqui-core/internal/repo"
	"zuqui-core/internal/service/auth"
	"zuqui-core/internal/service/email"
	"zuqui-core/internal/service/quiz"
)

type App struct {
	*fiber.App

	repo  *repo.Repo
	auth  auth.Service
	email email.Service
	quiz  *quiz.Service
}

func New(
	repo *repo.Repo,
	auth auth.Service,
	email email.Service,
	quiz *quiz.Service,
) *App {
	server := &App{
		App: fiber.New(fiber.Config{
			ServerHeader: "zuqui",
			AppName:      "zuqui",
		}),

		repo:  repo,
		auth:  auth,
		email: email,
		quiz:  quiz,
	}

	server.registerRoutes()

	return server
}

func (s *App) registerRoutes() {
	s.App.Use(logger.New())
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	quizGroup := s.Group("/quiz", s.AuthMiddleware)
	quizGroup.Post("/language/binary/generate", s.LanguageGenerateBinary)
	quizGroup.Post("/language/scq/generate", s.LanguageGenerateSCQ)
	quizGroup.Post("/math/binary/generate", s.MathGenerateBinary)
	quizGroup.Post("/math/scq/generate", s.MathGenerateSCQ)

	authGroup := s.Group("/auth")
	authGroup.Post("/login/otp", s.AuthLoginOTP)
	authGroup.Post("/onboard", s.AuthOnboard)
	authGroup.Post("/token/refresh", s.AuthTokenRefresh)
	authGroup.Post("/webauthn/registration", s.WebAuthnRegistration)
	authGroup.Post("/webauthn/authentication", s.WebAuthnAuthentication)

	meGroup := s.Group("/me", s.AuthMiddleware)
	meGroup.Get("/profile", s.MeProfile)
	meGroup.Get("/usage", s.MeUsage)
}
