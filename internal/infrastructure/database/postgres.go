package database

import (
	"Proyectos_Go/internal/core/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	log.Println("🚀 Conexión a DigitalOcean PostgreSQL exitosa")
	log.Println("Ejecutando migraciones...")
	err = DB.AutoMigrate(&model.User{}, &model.MunicipalityInfo{}, &model.Category{}, &model.TouristDestination{})
	if err != nil {
		log.Fatal("Error en la migración de base de datos: ", err)
	}
	log.Println("✅ Migraciones completadas")
}
