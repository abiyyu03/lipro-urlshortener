package handler

import (
	"time"
	"url-shortener/config"
	"url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

type StoreUrl struct{}

func (s *StoreUrl) CreateCustomUrl(c *fiber.Ctx) error {
	var requests *models.Request

	if err := c.BodyParser(&requests); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	customShort := requests.CustomShort
	expiry := requests.Expiry
	url := requests.OriginalUrl

	newUrl := models.Request{
		CustomShort: customShort,
		OriginalUrl: url,
		Expiry:      expiry * 3600 * time.Second,
	}
	// return c.JSON(newUrl)
	config.Db.Create(&newUrl)
	return c.Status(201).JSON(fiber.Map{
		"code":    201,
		"message": "Data created successfully",
		"data":    &newUrl,
	})
}
