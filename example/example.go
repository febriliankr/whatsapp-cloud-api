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

	myPhoneId := os.Getenv("MY_PHONE_ID")
	token := os.Getenv("TOKEN")

	wa := whatsapp.NewWhatsapp(token, myPhoneId)

	parameters := wa.GenerateTemplateParameters("text", "48884")

	components := wa.TemplateComponent("body", parameters)

	request := wa.CreateSendTemplateRequest(receiverPhoneNumber, whatsapp.TemplateVerifyPhoneNumberID, whatsapp.Indonesian, components)

	res, err := wa.SendWithTemplate(request)

	if err != nil {
		panic(err)
	}

	log.Println("res", res)

}
