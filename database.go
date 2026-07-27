package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// getEnv devuelve el valor de la variable de entorno o un valor por defecto.
func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}

func ConnectDatabase() {
	dbUser := getEnv("DB_USER", "rubi")
	dbPass := getEnv("DB_PASS", "1234")
	dbHost := getEnv("DB_HOST", "54.158.229.20")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "Veterinary")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	DB = database
}
