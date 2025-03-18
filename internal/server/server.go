package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"zuqui/internal/repo"
	authRoutes "zuqui/internal/server/routes/auth"
	quizRoutes "zuqui/internal/server/routes/quiz"
	userRoutes "zuqui/internal/server/routes/user"
	"zuqui/internal/service/auth"
	"zuqui/internal/service/email"
	"zuqui/internal/service/quiz"
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

func (app *App) registerRoutes() {
	app.App.Use(logger.New())
	app.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	authRoutes.RegisterRoutes(app, app.repo.User, app.auth, app.email, app.quiz)
	userRoutes.RegisterRoutes(app, app.auth)
	quizRoutes.RegisterRoutes(app, app.auth, app.quiz)
}
