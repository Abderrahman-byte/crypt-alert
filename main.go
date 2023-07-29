package main

import (
	"log"
	"os"

	"github.com/Abderrahman-byte/crypto-alert/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New()

	app.Post("/twilio/webhook", func(c *fiber.Ctx) error {
		message := twilioApi.ApiV2010Message{}
		err := c.BodyParser(&message)

		if err != nil {
			log.Fatal(err)
			return c.JSON(fiber.Map{"success": true})
		}

        body := *message.Body
        from := *message.From

        log.Printf("received message from %s : %s\n", from, body)
		utils.SendWhatsappMessage(from, "You said: "+ body)

		return c.JSON(fiber.Map{"success": true})
	})

	log.Fatal(app.Listen(os.Getenv("HTTP_LISTEN_ADDRESS")))
}
