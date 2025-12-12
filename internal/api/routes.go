package api

import (
	"Proyectos_Go/internal/api/handlers"
	"Proyectos_Go/internal/api/middleware"
	"Proyectos_Go/internal/core/service"
	"Proyectos_Go/internal/infrastructure/database"
	"Proyectos_Go/internal/infrastructure/repository"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://143.198.236.81"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := database.DB

	userRepo := repository.NewUserRepository(db)
	tourismRepo := repository.NewTourismRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)

	authService := service.NewAuthService(userRepo)
	tourismService := service.NewTourismService(tourismRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	authHandler := handlers.NewAuthHandler(authService)
	tourismHandler := handlers.NewTourismHandler(tourismService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		api.GET("/categories", categoryHandler.GetAll)

		api.GET("/tourism", tourismHandler.GetAll)
		api.GET("/tourism/:id", tourismHandler.GetByID)

		protected := api.Group("/")

		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/me", func(c *gin.Context) {
				userID, _ := c.Get("userID")
				role, _ := c.Get("role")

				c.JSON(200, gin.H{
					"message": "¡Acceso autorizado por Cookie!",
					"user_id": userID,
					"role":    role,
				})
			})

			protected.POST("/tourism", tourismHandler.Create)
			protected.PUT("/tourism/:id", tourismHandler.Update)
			protected.DELETE("/tourism/:id", tourismHandler.Delete)

			protected.POST("/categories", categoryHandler.Create)
		}
	}
}
