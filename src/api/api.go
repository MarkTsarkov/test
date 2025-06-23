package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/marktsarkov/test/service"
	"strconv"
)

func saveClick(service service.Iservice) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bannerIDStr := c.Params("bannerID")
		bannerID, err := strconv.Atoi(bannerIDStr)
		if err != nil {
			return c.Status(400).SendString(fmt.Sprintf("invalid banner id: %d, bannerIDStr:%s", bannerID, bannerIDStr))
		}
		err = service.SaveClick(bannerID)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString("ok")
	}
}
