// cmd/main.go
package main

import (
	"Proyectos_Go/config"
	"Proyectos_Go/internal/api"
	"Proyectos_Go/internal/infrastructure/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.ConnectDB()

	router := gin.Default()
	api.SetupRoutes(router)

	log.Printf("Servidor iniciado en puerto %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
