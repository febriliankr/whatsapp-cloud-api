package whatsapp

type Whatsapp struct {
	Token         string
	APIVersion    string
	PhoneNumberID string
	Language      TemplateLanguage
}

// Create new Whatsapp instance with v14.0 version and Indonesian as default language
func NewWhatsapp(token string, phoneNumberID string) *Whatsapp {
	return &Whatsapp{
		Language:      LanguageIndonesian,
		Token:         token,
		APIVersion:    "v14.0",
		PhoneNumberID: phoneNumberID,
	}
}

func (wa *Whatsapp) SendText(toPhoneNumber string, text string) (res map[string]interface{}, err error) {

	request := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                toPhoneNumber,
		"type":              "text",
		"text": map[string]string{
			"body": string(text),
		},
	}
	return wa.sendMessage(request)

}

// Sending the whatsapp message
//  1. `templateName` and `language` can be found in `whatsapp/constants.go` or in your template list dashboard https://business.facebook.com/wa/manage/message-templates
//  2. `receiverPhoneNumber` is the phone number that will receive the message (eg: 62852000000)
//  3. `components` is the parameters that will be sent to the receiver (eg: "999999" for OTP), can be empty if your template has no components
//     components parameter can be empty/nil if you don not want to send any parameters
func (wa *Whatsapp) SendWithTemplate(receiverPhoneNumber string, templateName string, components []Components) (res map[string]interface{}, err error) {

	request := wa.createSendWithTemplateRequest(receiverPhoneNumber, templateName, wa.Language, components)

	return wa.sendMessage(request)
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
func (wa *Whatsapp) createSendWithTemplateRequest(receiverPhoneNumber string, templateName string, language TemplateLanguage, components []Components) (res SendWithTemplateRequest) {
	return SendWithTemplateRequest{
		MessagingProduct: "whatsapp",
		To:               receiverPhoneNumber,
		Type:             "template",
		Template: Template{
			Name:     templateName,
			Language: LanguageIndonesian,
			// Components can be empty if you don't want to send any parameters
			Components: components,
		},
	}
}
