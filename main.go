package main

import (
	"fmt"
	"log"
	"manage-system-server/api/router"

	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbValue := DB{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     5432,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbValue.Host, dbValue.User, dbValue.Password, dbValue.Name, dbValue.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Connected to the database successfully:", db)

	router.SetUp(db, e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
