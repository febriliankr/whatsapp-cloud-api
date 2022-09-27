package whatsapp

// Struct for the response, currently not used because I'm afraid of breaking the wrapper if the response changes
type SendWithTemplateResponse struct {
	Contacts         []Contacts `json:"contacts"`
	Messages         []Messages `json:"messages"`
	MessagingProduct string     `json:"messaging_product"`
}

type Contacts struct {
	Input string `json:"input"`
	WaID  string `json:"wa_id"`
}

type Messages struct {
	ID string `json:"id"`
}
