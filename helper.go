package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Media struct {
	MessagingProduct string `json:"messaging_product"`
	URL              string `json:"url"`
	MIMEType         string `json:"mime_type"`
	SHA256           string `json:"sha256"`
	FileSize         string `json:"file_size"`
	ID               string `json:"id"`
}

func (wa *Whatsapp) getMedia(mediaID string) (media Media, err error) {

	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s", wa.APIVersion, mediaID)

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return media, err
	}

	req.Header.Set("Authorization", "Bearer "+wa.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return media, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&media)

	if err != nil {
		return media, err
	}
	return media, err
}

func (wa *Whatsapp) UploadMedia(filepath string) (id string, err error) {
	endpoint := fmt.Sprintf("https://graph.facebook.com/%s/%s/media", wa.APIVersion, wa.PhoneNumberID)

	// read filepath and get base64
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return id, err
	}

	// add messaging_product="whatsapp" to form data body
	resp, err := http.PostForm(endpoint, url.Values{
		"messaging_product": {"whatsapp"},
		"file":              {string(file)},
	})

	if err != nil {
		return id, err
	}

	defer resp.Body.Close()

	// check resp http status
	if resp.StatusCode != 200 {
		err := parseHTTPError(resp.Body)
		return id, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return id, err
	}
	fmt.Printf("%s\n", string(body))
	var res map[string]string
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&res)
	if err != nil {
		return id, err
	}
	id = res["id"]
	return id, err
}

// Http post request to send the message
func (wa *Whatsapp) sendMessage(request any) (res map[string]interface{}, err error) {

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return res, err
	}

	body := bytes.NewReader(jsonRequest)

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

	if resp.StatusCode != 200 {
		err := parseHTTPError(resp.Body)
		return res, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.Unmarshal(bodyBytes, &req)

	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)

	if err != nil {
		return res, err
	}
err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&res)

	if err != nil {
		return res, err
	}

	return res, err
}
