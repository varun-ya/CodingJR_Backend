package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"your_project/models"
)

type NoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateNote(db *gorm.DB)  fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		var input NoteInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		if input.Title == "" || input.Content == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title and content required"})
		}
		note := models.Note{
			Title: input.Title,
			Content: input.Content,
			UserID: userID,
		}
		if err := db.Create(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create note"})
		}
		return c.Status(fiber.StatusCreated).JSON(note)
	}
}

func ListNotes(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		var notes []models.Note
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		q := c.Query("q", "")
		offset := (page - 1) * limit
		query := db.Where("user_id = ?", userID)
		if q != "" {
			query = query.Where("title LIKE ? OR content LIKE ?", "%"+q+"%", "%"+q+"%")
		}
		if err := query.Offset(offset).Limit(limit).Order("created_at desc").Find(&notes).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch notes"})
		}
		return c.JSON(notes)
	}
}

func GetNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")
		var note models.Note
		if err := db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}
		return c.JSON(note)
	}
}

func UpdateNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")
		var note models.Note
		if err := db.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}
		var input NoteInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		note.Title = input.Title
		note.Content = input.Content
		if err := db.Save(&note).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update note"})
		}
		return c.JSON(note)
	}
}

func DeleteNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		id := c.Params("id")
		if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete note"})
		}
		return c.JSON(fiber.Map{"message": "Note deleted"})
	}
} 