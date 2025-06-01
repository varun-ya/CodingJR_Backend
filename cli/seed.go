package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/varun-ya/CodingJR_backend.git/models" 
	"github.com/varun-ya/CodingJR_backend.git/utils"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection failed")
	}


	db.AutoMigrate(&models.User{}, &models.Note{})

	
	users := []models.User{}
	for i := 1; i <= 4; i++ {
		hash, _ := utils.HashPassword("password123")
		user := models.User{
			Name:     fmt.Sprintf("User%d", i),
			Email:    fmt.Sprintf("user%d@example.com", i),
			Password: hash,
		}
		db.Create(&user)
		users = append(users, user)
	}


	for _, user := range users {
		for j := 1; j <= 3; j++ {
			note := models.Note{
				Title:   fmt.Sprintf("Note %d for %s", j, user.Name),
				Content: fmt.Sprintf("This is note %d for %s", j, user.Name),
				UserID:  user.ID,
			}
			db.Create(&note)
		}
	}

	fmt.Println(" Seeded 4 users and 12 notes successfully.")
}
