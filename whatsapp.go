package whatsapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Whatsapp struct {
	Token         string
	APIVersion    string
	PhoneNumberID string
}

func NewWhatsapp(token string, phoneNumberID string) *Whatsapp {
	return &Whatsapp{
		Token:         token,
		APIVersion:    "v14.0",
		PhoneNumberID: phoneNumberID,
	}
}

// Sending the whatsapp message
func (wa *Whatsapp) SendWithTemplate(request SendTemplateRequest) (*http.Response, error) {
	marshaledJSON, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	reqString := string(marshaledJSON)

	body := strings.NewReader(reqString)

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", wa.APIVersion, wa.PhoneNumberID)
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+wa.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, err
}

// Parameter if exists
// Using GenerateTemplateParameters("text", "Good evening, Febrilian")
// Generates this json for the "component" field in the Whatsapp CloudAPI body request:
// `{ "type": "text", "text": "Good evening, Febrilian" }“
// you can append between parameters if you have multiple `parameterType`, for example:
// ```
// parameters := wa.GenerateTemplateParameters("text", "48884")
// mediaParameters := wa.GenerateTemplateParameters("media", "media_url")
// parameters = append(parameters, mediaParameters...)
// ```
func (wa *Whatsapp) GenerateTemplateParameters(parameterType string, args ...string) (res []TemplateParameters) {

	if parameterType == "" {
		parameterType = "text"
	}

	for _, arg := range args {
		res = append(res, TemplateParameters{
			Type: parameterType,
			Text: arg,
		})
	}
	return
}

// Forging component parameter for the request. Example usage:
// `component:= TemplateComponent("body", TemplateParametersText("text", "999999"))“
// Will generate a struct for the `Component` field, will fill the template parameter (eg: `{{1}}`)
func (wa *Whatsapp) TemplateComponent(componentType string, args ...[]TemplateParameters) (components []Components) {

	if componentType == "" {
		componentType = "body"
	}

	for _, arg := range args {
		components = append(components, Components{
			Type:       componentType,
			Parameters: arg,
		})
	}

	return components
}

// Forge request with receiverPhoneNumber, templateName, Language, and components []Components
//  1. `templateName` and `language` can be found in `whatsapp/constants.go` or in your template list dashboard (https://business.facebook.com/wa/manage/message-templates/?business_id=886970291828176&waba_id=115257484519594)
//  2. `receiverPhoneNumber` is the phone number that will receive the message (eg: 62852000000)
//  3. `components` is the parameters that will be sent to the receiver (eg: "999999" for OTP), can be empty if your template has no components
//     components parameter can be empty/nil if you don't want to send any parameters
func (wa *Whatsapp) CreateSendTemplateRequest(receiverPhoneNumber string, templateName string, language TemplateLanguage, components []Components) (res SendTemplateRequest) {
	return SendTemplateRequest{
		MessagingProduct: "whatsapp",
		To:               receiverPhoneNumber,
		Type:             "template",
		Template: Template{
			Name:     TemplateVerifyPhoneNumberID,
			Language: Indonesian,
			// Components can be empty if you don't want to send any parameters
			Components: components,
		},
	}
}
