package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
	"strings"
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
			return c.JSON(&fiber.Map{"success": true})
		}

		log.Printf("received message from %s\n", message.From)

		SendWhatsappMessage(message.From, "You said: "+message.Body)
		return c.JSON(&fiber.Map{"success": true})
	})

	app.Listen(":3000")
}

func SendWhatsappMessage(to string, body string) {
	client := twilio.NewRestClient()

	params := twilioApi.CreateMessageParams{}
	params.SetFrom("whatsapp:" + os.Getenv("TWILIO_PHONE"))
	params.SetBody(body)

	if strings.HasPrefix(to, "whatsapp:") {
		params.SetTo(to)
	} else {
		params.SetTo("whatsapp:" + to)
	}

	_, err := client.Api.CreateMessage(&params)

	if err != nil {
		log.Fatal(err)
	}
}
