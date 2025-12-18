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
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
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
	userService := service.NewUserService(userRepo)
	tourismService := service.NewTourismService(tourismRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
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
			protected.GET("/me", userHandler.GetMe)

			superAdmin := protected.Group("/admin")
			superAdmin.Use(middleware.RequireRole("superadmin"))
			{
				superAdmin.GET("/users", userHandler.GetAll)
				superAdmin.DELETE("/users/:id", userHandler.Delete)
				superAdmin.PATCH("/users/:id/role", userHandler.ChangeRole)
			}

			contentManagers := protected.Group("/")
			contentManagers.Use(middleware.RequireRole("superadmin", "admin"))
			{
				contentManagers.POST("/tourism", tourismHandler.Create)
				contentManagers.PUT("/tourism/:id", tourismHandler.Update)
				contentManagers.DELETE("/tourism/:id", tourismHandler.Delete)

				contentManagers.POST("/categories", categoryHandler.Create)
			}
		}
	}
}
