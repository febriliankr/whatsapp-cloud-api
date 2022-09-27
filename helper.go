package whatsapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (wa *Whatsapp) sendMessage(request any) (res map[string]interface{}, err error) {

	marshaledJSON, err := json.Marshal(request)
	if err != nil {
		return res, err
	}
	reqString := string(marshaledJSON)

	body := strings.NewReader(reqString)

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", wa.APIVersion, wa.PhoneNumberID)
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return res, err
	}
	req.Header.Set("Authorization", "Bearer "+wa.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return res, err
	}

	return res, err
}
