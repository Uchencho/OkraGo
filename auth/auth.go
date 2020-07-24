package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	if u.token == "" || u.baseurl == "" {
		panic("Token and Base url is needed to call this Function")
	}
	return u
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

// ByID fetch authentication info using the id of the authentication record.
func (w Initializer) ByID(i string) (scode int, body string, err error) {

	reqBody, err := json.Marshal(map[string]string{
		"id": i,
	})
	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + w.token
	endpoint := w.baseurl + "auth/getById"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearer)

	scode, body, err = postRequest(req)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

// ByOptions fetch authentication info using the options metadata you provided when setting up the widget.
func (w Initializer) ByOptions(page string, limit string, firstname string, lastname string) (scode int, body string, err error) {

	reqBody, err := json.Marshal(map[string]string{
		"page":  page,
		"limit": limit,
		// "options": {"first_name": firstname, "last_name": lastname},
	})
	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + w.token
	endpoint := w.baseurl + "auth/byOptions"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearer)

	scode, body, err = postRequest(req)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

// ByCustomer fetch authentication info using the customer id
func (w Initializer) ByCustomer(customerID string) (scode int, body string, err error) {

	reqBody, err := json.Marshal(map[string]string{
		"customer": customerID,
	})
	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + w.token
	endpoint := w.baseurl + "auth/getByCustomer"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	req.Header.Add("Authorization", bearer)

	scode, body, err = postRequest(req)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
