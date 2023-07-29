package main

import (
	"github.com/Abderrahman-byte/crypto-alert/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type TwillioMessage struct {
	Body string
	From string
}

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New()

	app.Post("/twilio/webhook", func(c *fiber.Ctx) error {
		message := TwillioMessage{}
		err := c.BodyParser(&message)

		if err != nil {
			log.Fatal(err)
			return c.JSON(fiber.Map{"success": true})
		}

		log.Printf("received message from %s\n", message.From)

		utils.SendWhatsappMessage(message.From, "You said: "+message.Body)
		return c.JSON(fiber.Map{"success": true})
	})

	log.Fatal(app.Listen(os.Getenv("HTTP_LISTEN_ADDRESS")))
}
