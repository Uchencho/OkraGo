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

type initializer struct {
	token   string
	baseurl string
}

func (w initializer) byID(i string) (scode int, body string, err error) {

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
