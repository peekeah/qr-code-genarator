package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

type ReqBody struct {
	Text string `json:"text"`
}

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {

		body := new(ReqBody)

		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(403)
		}

		if body.Text == "" {
			c.Status(403)
			return c.SendString("text field is required")
		}

		var png []byte
		png, err := qrcode.Encode(body.Text, qrcode.Medium, 256)

		// to store in file
		// err := qrcode.WriteFile(body.Text, qrcode.Medium, 256, "qr.png")
		if err != nil {
			return c.SendStatus(500)
		}
		c.Set("Content-type", "image/jpeg; charset=UTF-8")
		return c.Send(png)
	})
	app.Listen(":8000")
}
