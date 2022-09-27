package whatsapp

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// func TestWhatsapp_SendWithTemplate(t *testing.T) {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	WHATSAPP_PHONE_ID := os.Getenv("WHATSAPP_PHONE_ID")
// 	WHATSAPP_TOKEN := os.Getenv("WHATSAPP_TOKEN")
// 	RECIPIENT_NUMBER := os.Getenv("RECIPIENT_NUMBER")

// 	wa := NewWhatsapp(WHATSAPP_TOKEN, WHATSAPP_PHONE_ID)

// 	parameters := wa.GenerateTemplateParameters("text", "48884")

// 	components := wa.TemplateComponent("body", parameters)

// 	_, err = wa.SendWithTemplate(RECIPIENT_NUMBER, TemplateVerifyPhoneNumberID, components)

// 	if err != nil {
// 		t.Error(err)
// 	}

// }

func TestWhatsapp_SendText(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	WHATSAPP_PHONE_ID := os.Getenv("WHATSAPP_PHONE_ID")
	WHATSAPP_TOKEN := os.Getenv("WHATSAPP_TOKEN")

	wa := NewWhatsapp(WHATSAPP_TOKEN, WHATSAPP_PHONE_ID)

	res, err := wa.SendText("__", "testing")

	if err != nil {
		t.Error(err)
	}
	log.Println("res", res)

}

// func TestWhatsapp_UploadMedia(t *testing.T) {

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	WHATSAPP_PHONE_ID := os.Getenv("WHATSAPP_PHONE_ID")
// 	WHATSAPP_TOKEN := os.Getenv("WHATSAPP_TOKEN")

// 	wa := NewWhatsapp(WHATSAPP_TOKEN, WHATSAPP_PHONE_ID)

// 	_, err = wa.UploadMedia("./assets/meta-analysis-email-header.jpg")

// 	if err != nil {
// 		t.Error(err)
// 	}

// }
