package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v2"
	"google.golang.org/api/option"

	"zuqui-core/internal"
	"zuqui-core/internal/repo"
	"zuqui-core/internal/server"
	"zuqui-core/internal/service/auth"
	"zuqui-core/internal/service/email"
	"zuqui-core/internal/service/quiz"
)

func gracefulShutdown(fiberServer *server.App, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	resend, ai, redis, db, cleanup := initServices()
	defer cleanup()

	server := server.New(
		repo.New(db),
		auth.New(redis),
		email.New(resend),
		quiz.New(ai),
	)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	go func() {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}

func initServices() (*resend.Client, *genai.Client, *redis.Client, *pgxpool.Pool, func()) {
	resend := resend.NewClient(internal.Env.RESEND_API_KEY)

	ai, err := genai.NewClient(context.Background(), option.WithAPIKey(internal.Env.GEMINI_API_KEY))
	if err != nil {
		fmt.Println("Error creating genai client:", err)
		os.Exit(1)
	}

	redisOpt, err := redis.ParseURL(internal.Env.UPSTASH_REDIS_URI)
	if err != nil {
		fmt.Println("Error parsing redis options:", err)
		os.Exit(1)
	}
	redis := redis.NewClient(redisOpt)

	db, err := pgxpool.New(context.Background(), internal.Env.DATABASE_URL)
	if err != nil {
		fmt.Println("Error connecting database:", err)
		os.Exit(1)
	}

	return resend, ai, redis, db, func() {
		ai.Close()
		redis.Close()
	}
}
