package utils

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
	"strings"
)

func SendWhatsappMessage(to string, body string) error {
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

	return err
}
