package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	urlByID = "/auth/getById"
)

// Initializer struct
type Initializer struct {
	token   string
	baseurl string
}

// New creates a new struct
func New(t string, b string) Initializer {
	u := Initializer{
		token:   t,
		baseurl: b,
	}
	return u
}

func (w Initializer) byID(i string) (scode int, body string, err error) {

	if w.token == "" || w.baseurl == "" {
		panic("Token and Base url is needed to call this method")
	}

	reqBody, err := json.Marshal(map[string]string{
		"id": i,
	})
	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + w.token
	req, err := http.NewRequest("POST", urlByID, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearer)

	scode, body, err = postRequest(req)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func postRequest(req *http.Request) (scode int, body string, err error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	scode = resp.StatusCode
	body = string(bod)
	return
}

// RetrieveAuth retrieves authentication of a user
func (w Initializer) RetrieveAuth() (sCode int, body string, err error) {
	var bearer = "Bearer " + w.token
	const url = "products/auths"
	endpoint := w.baseurl + url
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", bearer)

	sCode, body, err = postRequest(req)
	if err != nil {
		log.Fatalln(err)
	}
	return

}
