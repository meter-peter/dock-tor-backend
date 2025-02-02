package main

import (
	"clinic-management/api/handlers"
	"clinic-management/api/middleware"
	"clinic-management/api/routes"
	"clinic-management/config"
	_ "clinic-management/docs"
	"clinic-management/internal/models"
	"clinic-management/internal/repositories"
	"clinic-management/internal/services"
	"clinic-management/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

// @title Clinic Management API
// @version 1.0
// @description This is a clinic management server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load configuration
	cfg := config.GetConfig()

	// Initialize database
	config.ConnectDB()
	db := config.DB

	// Auto-migrate models
	err := db.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Appointment{},
		&models.MedicalHistory{},
		&models.HistoryEntry{},
		&models.DoctorAvailability{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize MinIO
	utils.InitMinIO()

	// Initialize repositories
	userRepo := repositories.NewUserRepository()
	patientRepo := repositories.NewPatientRepository()
	appointmentRepo := repositories.NewAppointmentRepository()
	availabilityRepo := repositories.NewDoctorAvailabilityRepository() // Added
	// Initialize services
	authService := services.NewAuthService(userRepo)
	patientService := services.NewPatientService(patientRepo)
	appointmentService := services.NewAppointmentService(
		appointmentRepo,
		availabilityRepo, // Added
	)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	patientHandler := handlers.NewPatientHandler(patientService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	router := gin.Default()
	router.Use(
		middleware.CORS(),
		middleware.Logger(),
		middleware.Recovery(),
	)

	// Configure routes
	routes.ConfigureRoutes(
		router,
		authHandler,
		patientHandler,
		appointmentHandler,
	)

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("💥 Failed to start server: %v", err)
	}
}
