package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("File upload error")
		}

		filePath := fmt.Sprintf("./%s", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file")
		}

		client := gosseract.NewClient()
		defer client.Close()
		client.SetImage(filePath)
		text, err := client.Text()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to perform OCR")
		}

		return c.JSON(fiber.Map{
			"text": text,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
