package routes

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"zuqui-core/internal/server"
	"zuqui-core/internal/server/routes/auth"
	"zuqui-core/internal/server/routes/me"
	"zuqui-core/internal/server/routes/quiz/language"
	"zuqui-core/internal/server/routes/quiz/math"
)

func RegisterFiberRoutes(s *server.FiberServer) {
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	quizGroup := s.Group("/quiz")
	quizGroup.Post("/language/binary/generate", language.GenerateBinary)
	quizGroup.Post("/language/scq/generate", language.GenerateSCQ)
	quizGroup.Post("/math/binary/generate", math.GenerateBinary)
	quizGroup.Post("/math/scq/generate", math.GenerateSCQ)

	authGroup := s.Group("/auth")
	authGroup.Post("/login/otp", auth.LoginOTP)
	authGroup.Post("/webauthn/registration", auth.WebAuthnRegistration)
	authGroup.Post("/webauthn/authentication", auth.WebAuthnAuthentication)

	meGroup := s.Group("/me")
	meGroup.Get("/profile", me.Profile)
	meGroup.Get("/usage", me.Usage)
}
