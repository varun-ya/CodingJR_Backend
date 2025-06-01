package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"your_project/models"
	"your_project/routes"
	"your_project/utils"
)

func main() {
	fmt.Println("Loaded ENV before godotenv.Load(): DB_HOST =", os.Getenv("DB_HOST"), ", PORT =", os.Getenv("PORT"))

	if _, err := os.Stat(".env"); err == nil {
		for _, e := range os.Environ() {
			fmt.Println("ENV:", e)
		}
		err := godotenv.Load()
		if err != nil {
			log.Println("Failed to load .env file:", err)
		} else {
			fmt.Println("Loaded .env file")
		}
	} else {
		log.Println("No .env file found or unable to load. Assuming running in container...")
	}

	fmt.Println("Loaded ENV after godotenv.Load(): DB_HOST =", os.Getenv("DB_HOST"), ", PORT =", os.Getenv("PORT"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{}, &models.Note{})

	utils.SetJWTSecret(os.Getenv("JWT_SECRET"))

	app := fiber.New()
	routes.Setup(app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running")
	})

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("No PORT set. Defaulting to 3000")
		port = "3000"
	}
	if port == "3306" {
		log.Fatal("PORT is set to 3306 (MySQL port). Set PORT=3000 in your environment variables")
	}

	log.Println("Starting server on port", port)
	log.Fatal(app.Listen(":" + port))
}
