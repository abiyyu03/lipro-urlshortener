package handler

import (
	"url-shortener/config"
	"url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

type ResolveUrl struct{}

func (*ResolveUrl) Resolve(c *fiber.Ctx) error {
	//get dari params
	// cocokan ke database
	// jika true redirect
	var requests *models.Request
	url := c.Params("short_url")

	config.Db.Where("custom_short = ?", url).First(&requests)
	// return c.JSON(requests.OriginalUrl)
	return c.Redirect(requests.OriginalUrl, 302)
}
