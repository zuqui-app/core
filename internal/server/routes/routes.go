package routes

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"zuqui-core/internal/server"
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

	s.App.Post("/quiz/language/binary/generate", language.GenerateBinary)
	s.App.Post("/quiz/language/scq/generate", language.GenerateSCQ)
	s.App.Post("/quiz/math/binary/generate", math.GenerateBinary)
	s.App.Post("/quiz/math/scq/generate", math.GenerateSCQ)
}
