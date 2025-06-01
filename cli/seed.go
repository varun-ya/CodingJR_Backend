package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"your_project/models"
	"your_project/utils"
)

func main() {
	godotenv.Load("../.env")
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
	users := []models.User{}
	for i := 1; i <= 5; i++ {
		hash, _ := utils.HashPassword("password123")
		user := models.User{
			Name: fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
			Password: hash,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}
		db.Create(&user)
		users = append(users, user)
	}
	for _, user := range users {
		for j := 1; j <= 3; j++ {
			note := models.Note{
				Title: fmt.Sprintf("Note %d for %s", j, user.Name),
				Content: fmt.Sprintf("This is note %d for %s", j, user.Name),
				UserID: user.ID,
				CreatedAt: time.Now().Unix() + int64(rand.Intn(1000)),
				UpdatedAt: time.Now().Unix() + int64(rand.Intn(1000)),
			}
			db.Create(&note)
		}
	}
	fmt.Println("Seeded 4 users and 12 notes.")
} 