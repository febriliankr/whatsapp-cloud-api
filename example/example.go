package example

import (
	"log"
	"os"

	"github.com/febriliankr/whatsapp-cloud-api"
	"github.com/joho/godotenv"
)

var receiverPhoneNumber = "62810000000"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myPhoneId := os.Getenv("WHATSAPP_PHONE_ID")
	token := os.Getenv("WHATSAPP_TOKEN")

	wa := whatsapp.NewWhatsapp(token, myPhoneId)

	parameters := wa.GenerateTemplateParameters("text", "48884")

	components := wa.TemplateComponent("body", parameters)

	res, err := wa.SendWithTemplate(receiverPhoneNumber, whatsapp.TemplateVerifyPhoneNumberID, components)

	if err != nil {
		panic(err)
	}

	log.Println("res", res)

}
