// api/routes/routes.go

package routes

import (
	"clinic-management/api/handlers"
	"clinic-management/api/middleware"
	"clinic-management/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(router *gin.Engine, services *services.Services) {
	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	router.GET("/health", handlers.HealthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			authHandler := handlers.NewAuthHandler(services.AuthService)
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Patient routes
			patients := protected.Group("/patients")
			{
				patientHandler := handlers.NewPatientHandler(services.PatientService)
				patients.POST("", patientHandler.CreatePatient)
				patients.GET("", patientHandler.ListPatients)
				patients.GET("/:id", patientHandler.GetPatient)
				patients.PUT("/:id", patientHandler.UpdatePatient)
				patients.DELETE("/:id", patientHandler.DeletePatient)
				patients.POST("/bulk", patientHandler.BulkUploadPatients)
			}

			// Appointment routes
			appointments := protected.Group("/appointments")
			{
				appointmentHandler := handlers.NewAppointmentHandler(services.AppointmentService)
				appointments.POST("", appointmentHandler.CreateAppointment)
				appointments.GET("", appointmentHandler.ListAppointments)
				appointments.GET("/:id", appointmentHandler.GetAppointment)
				appointments.PUT("/:id", appointmentHandler.UpdateAppointment)
				appointments.DELETE("/:id", appointmentHandler.DeleteAppointment)
			}

			// Doctor routes
			doctors := protected.Group("/doctors")
			{
				doctorHandler := handlers.NewDoctorHandler(services.DoctorService)
				doctors.POST("", doctorHandler.CreateDoctor)
				doctors.GET("", doctorHandler.ListDoctors)
				doctors.GET("/:id", doctorHandler.GetDoctor)
				doctors.PUT("/:id", doctorHandler.UpdateDoctor)
				doctors.DELETE("/:id", doctorHandler.DeleteDoctor)
				doctors.POST("/:id/availabilities", doctorHandler.CreateAvailability)
				doctors.GET("/:id/availabilities", doctorHandler.GetAvailabilities)
				doctors.DELETE("/availabilities/:id", doctorHandler.DeleteAvailability)
			}

			// Medical History routes
			history := protected.Group("/history")
			{
				historyHandler := handlers.NewHistoryHandler(services.HistoryService)
				history.POST("", historyHandler.AddHistoryEntry)
				history.GET("/:medicalHistoryID", historyHandler.GetPatientHistory)
				history.PUT("/latest", historyHandler.UpdateLatestEntry)
				history.DELETE("/:medicalHistoryID/latest", historyHandler.DeleteLatestEntry)
				history.GET("/search", historyHandler.SearchHistory)
			}
		}
	}
}
