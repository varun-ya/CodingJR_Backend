package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"your_project/handlers"
	"your_project/middleware"
)

func Setup(app *fiber.App,  db *gorm.DB) {
	app.Post("/register",  handlers.Register(db))
	app.Post("/login",  handlers.Login(db))

	notes := app.Group("/notes", middleware.JWTProtected())
	notes.Post("/", handlers.CreateNote(db))
	notes.Get("/", handlers.ListNotes(db))
	notes.Get(":id", handlers.GetNote(db))
	notes.Put(":id", handlers.UpdateNote(db))
	notes.Delete(":id", handlers.DeleteNote(db))
} 