package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/config"
	"github.com/seu-usuario/solis-backend/dataprovider/database/migration"
	"github.com/seu-usuario/solis-backend/entrypoint/http/routes"
	"github.com/seu-usuario/solis-backend/seeders"
)

func main() {
	// 1. Parse command line flags
	migrate := flag.Bool("migrate", false, "Run database migrations")
	seed := flag.Bool("seed", false, "Run database seeders")
	flag.Parse()

	// 2. Load configuration from .env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// 3. Validate configuration
	if err := cfg.Validate(); err != nil {
		log.Fatalf("❌ Invalid configuration: %v", err)
	}

	log.Printf("📋 Configuration loaded (Environment: %s)", cfg.Server.Environment)

	// 4. Create dependency injection container
	container, err := config.NewContainer(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to create container: %v", err)
	}
	defer container.Close()

	log.Println("✅ Dependency container initialized")

	// 5. Run database migrations if flag is set
	if *migrate {
		log.Println("🔄 Running database migrations...")
		if err := migration.RunMigrations(container.DB); err != nil {
			log.Fatalf("❌ Migration failed: %v", err)
		}
		log.Println("✅ Migrations completed successfully")
		return
	}

	// 6. Run database seeders if flag is set
	if *seed {
		log.Println("🌱 Running database seeders...")
		ctx := context.Background()
		if err := seeders.SeedAll(ctx, container.UserSeeder, container.KpiSeeder, container.SocialBenchmarkingSeeder, container.CooperativeSeeder); err != nil {
			log.Fatalf("❌ Seeders failed: %v", err)
		}
		log.Println("✅ Seeders completed successfully")
		return
	}

	// 7. Create Fiber application
	app := fiber.New(fiber.Config{
		AppName:               "Solis MKT API",
		DisableStartupMessage: false,
		ErrorHandler:          customErrorHandler,
	})

	// 8. Setup routes
	routes.SetupRoutes(
		app,
		container.GetControllers(),
		container.GetMiddlewares(),
	)

	log.Println("✅ Routes configured")

	// 9. Start server with graceful shutdown
	startServer(app, cfg)
}

// customErrorHandler handles errors globally and returns JSON responses
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error": message,
	})
}

// startServer starts the HTTP server with graceful shutdown support
func startServer(app *fiber.App, cfg *config.Config) {
	// Create channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		addr := fmt.Sprintf(":%d", cfg.Server.Port)
		log.Printf("🚀 Server starting on http://localhost:%d", cfg.Server.Port)
		log.Printf("📝 API Documentation: http://localhost:%d/api", cfg.Server.Port)
		log.Printf("💚 Health check: http://localhost:%d/health", cfg.Server.Port)

		if err := app.Listen(addr); err != nil {
			log.Fatalf("❌ Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	<-quit
	log.Println("🛑 Shutting down server...")

	// Graceful shutdown with timeout
	if err := app.Shutdown(); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server stopped gracefully")
}
