package controllers

import (

	"github.com/gofiber/fiber/v2"
)

/*
GetStores | @Desc: Get Stores by post json card_no |
@Method: POST |
@Route: "stores/list" |
@Auth: Public
*/
func HelloRegister(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"code": 200, "data": "Hello World", "success": true})
}
