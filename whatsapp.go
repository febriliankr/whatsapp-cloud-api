package whatsapp

type Whatsapp struct {
	Token         string
	APIVersion    string
	PhoneNumberID string
	Language      TemplateLanguage
}

// Create new Whatsapp instance with v14.0 version and English as default language
func NewWhatsapp(token string, phoneNumberID string) *Whatsapp {
	return &Whatsapp{
		Language:      LanguageEnglish,
		Token:         token,
		APIVersion:    "v14.0",
		PhoneNumberID: phoneNumberID,
	}
}

// Sending plain text message to a phone number that has messaged your WhatsApp Business account in the past 24 hours.
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

/*
Sending a templated whatsapp message. `templateName` can be found in your template list dashboard https://business.facebook.com/wa/manage/message-templates
`components` is the parameters that will be sent to the receiver (eg: "999999" for OTP), can be empty/nil if your template has no components.
*/
func (wa *Whatsapp) SendWithTemplate(toPhoneNumber string, templateName string, components []Components) (res map[string]interface{}, err error) {

	request := wa.createSendWithTemplateRequest(toPhoneNumber, templateName, wa.Language, components)

	return wa.sendMessage(request)
}

// Generating parameter if exists
/*
Using GenerateTemplateParameters("text", "Good evening, Febrilian")
Generates this json for the "component" field in the Whatsapp CloudAPI body request:
`{ "type": "text", "text": "Good evening, Febrilian" }“
you can append between parameters if you have multiple `parameterType`, for example:
*/

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
			Language: language,
			// Components can be empty if you don't want to send any parameters
			Components: components,
		},
	}
}
