package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"test-echo/internal/config"
	"test-echo/internal/handler"
	"test-echo/internal/repository"
	"test-echo/internal/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repository, service, and handler
	nasabahRepo := repository.NewNasabahRepository(db)
	nasabahService := service.NewNasabahService(nasabahRepo)
	nasabahHandler := handler.NewNasabahHandler(nasabahService)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/daftar", nasabahHandler.DaftarNasabah)
	e.POST("/tabung", nasabahHandler.Tabung)
	e.POST("/tarik", nasabahHandler.Tarik)
	e.GET("/saldo/:no_rekening", nasabahHandler.GetSaldo)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}