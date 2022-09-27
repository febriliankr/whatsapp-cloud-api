package whatsapp

type SendTemplateRequest struct {
	MessagingProduct string   `json:"messaging_product,omitempty"`
	To               string   `json:"to,omitempty"`
	Type             string   `json:"type,omitempty"`
	Template         Template `json:"template,omitempty"`
}
type TemplateLanguage struct {
	Code string `json:"code,omitempty"`
}
type TemplateParameters struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Components struct {
	Type       string               `json:"type,omitempty"`
	Parameters []TemplateParameters `json:"parameters,omitempty"`
}

type Template struct {
	Name       string           `json:"name,omitempty"`
	Language   TemplateLanguage `json:"language,omitempty"`
	Components []Components     `json:"components,omitempty"`
}
