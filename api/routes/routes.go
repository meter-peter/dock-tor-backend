package routes

import (
	"clinic-management/api/handlers"
	"clinic-management/api/middleware"
	"clinic-management/internal/repositories"
	"clinic-management/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	AuthHandler interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
		Logout(c *gin.Context)
	}

	PatientHandler interface {
		CreatePatient(c *gin.Context)
		GetPatient(c *gin.Context)
		UpdatePatient(c *gin.Context)
		DeletePatient(c *gin.Context)
		ListPatients(c *gin.Context)
		BulkUploadPatients(c *gin.Context)
	}

	AppointmentHandler interface {
		CreateAppointment(c *gin.Context)
		GetAppointment(c *gin.Context)
		UpdateAppointment(c *gin.Context)
		DeleteAppointment(c *gin.Context)
		ListAppointments(c *gin.Context)
	}

	HistoryHandler interface {
		AddHistoryEntry(c *gin.Context)
		GetHistory(c *gin.Context)
		UpdateHistoryEntry(c *gin.Context)
	}

	AvailabilityHandler interface {
		SetAvailability(c *gin.Context)
		GetAvailability(c *gin.Context)
	}
)

func ConfigureRoutes(
	router *gin.Engine,
	authHandler *handlers.AuthHandler,
	patientHandler PatientHandler,
	appointmentHandler AppointmentHandler,
) {
	// Swagger documentation (must be before any other middleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	// Public routes
	router.GET("/health", healthCheck)
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", middleware.Auth(), authHandler.Logout)
	}

	// API version 1
	v1 := router.Group("/api/v1")
	{
		// Authenticated routes
		authenticated := v1.Group("")
		authenticated.Use(middleware.Auth())
		{
			// Patient routes
			patients := authenticated.Group("/patients")
			patients.Use(middleware.RoleAuth("secretary", "doctor"))
			{
				patients.POST("", patientHandler.CreatePatient)
				patients.POST("/bulk", patientHandler.BulkUploadPatients)
				patients.GET("", patientHandler.ListPatients)
				patients.GET("/:id", patientHandler.GetPatient)
				patients.PUT("/:id", patientHandler.UpdatePatient)
				patients.DELETE("/:id", patientHandler.DeletePatient)
			}

			// Appointment routes
			appointments := authenticated.Group("/appointments")
			{
				appointments.POST("", middleware.RoleAuth("secretary", "doctor"), appointmentHandler.CreateAppointment)
				appointments.GET("", appointmentHandler.ListAppointments)
				appointments.GET("/:id", appointmentHandler.GetAppointment)
				appointments.PUT("/:id", middleware.RoleAuth("secretary", "doctor"), appointmentHandler.UpdateAppointment)
				appointments.DELETE("/:id", middleware.RoleAuth("secretary", "doctor"), appointmentHandler.DeleteAppointment)
			}

			doctorRepo := repositories.NewDoctorRepository()
			doctorService := services.NewDoctorService(doctorRepo)
			doctorHandler := handlers.NewDoctorHandler(doctorService)
			doctors := router.Group("/doctors")
			doctors.Use(middleware.Auth())
			{
				doctors.POST("", doctorHandler.CreateDoctor)
				doctors.GET("", doctorHandler.ListDoctors)
				doctors.GET("/:id", doctorHandler.GetDoctor)
				doctors.PUT("/:id", doctorHandler.UpdateDoctor)
				doctors.DELETE("/:id", doctorHandler.DeleteDoctor)
				doctors.POST("/:id/availabilities", doctorHandler.CreateAvailability)
				doctors.GET("/:id/availabilities", doctorHandler.GetAvailabilities)
				doctors.DELETE("/availabilities/:id", doctorHandler.DeleteAvailability)
			}
		}
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"version": "1.0.0",
	})
}
