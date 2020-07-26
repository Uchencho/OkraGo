package okra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Initializer struct
type Initializer struct {
	token   string
	baseurl string
}

type option struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type optionPayload struct {
	Page    string `json:"page"`
	Limit   string `json:"limit"`
	Options option `json:"options"`
}

type genPayload struct {
	Page       string `json:"page"`
	Limit      string `json:"limit"`
	CustomerID string `json:"customer"`
	From       string `json:"from"`
	To         string `json:"to"`
	BankID     string `json:"bank"`
	ID         string `json:"id"`
}

// NewOkra returns a struct that can be used to call all methods
func NewOkra(t, b string) Initializer {
	u := Initializer{
		token:   t,
		baseurl: b,
	}
	if u.token == "" || u.baseurl == "" {
		panic("Token and Base url is needed to call this Function")
	}
	return u
}

func postRequest(pl interface{}, url, token string) (body string, err error) {

	reqBody, err := json.Marshal(pl)
	if err != nil {
		return "Error", fmt.Errorf("error marshalling json: %w", err)
	}

	var bearer = "Bearer " + token
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "Error", fmt.Errorf("error making http call: %w", err)
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error", fmt.Errorf("error doing request: %w", err)
	}

	defer resp.Body.Close()

	bod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error", fmt.Errorf("error reading body: %w", err)
	}
	if resp.StatusCode != 200 {
		body = "Status code returned was: " + string(resp.StatusCode)
	} else {
		body = string(bod)
	}

	return
}

// RetrieveAuth retrieves authentication of a user
func (w Initializer) RetrieveAuth() (body string, err error) {

	endpoint := w.baseurl + "products/auths"
	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth token: %w", err)
	}
	return

}

// AuthByID fetches authentication info using the id of the authentication record.
func (w Initializer) AuthByID(page, limit, i string) (body string, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		ID:    i,
	}
	endpoint := w.baseurl + "auth/getById"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching auth using id: %w", err)
	}
	return
}

// AuthByOptions fetches authentication info using the options metadata you provided when setting up the widget.
func (w Initializer) AuthByOptions(page, limit, firstname, lastname string) (body string, err error) {

	pl := optionPayload{
		Page:  page,
		Limit: limit,
		Options: option{
			FirstName: firstname,
			LastName:  lastname,
		},
	}
	url := w.baseurl + "auth/getByOptions"

	body, err = postRequest(pl, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byoptions: %w", err)
	}
	return
}

// AuthByCustomer fetches authentication info using the customer id
func (w Initializer) AuthByCustomer(page, limit, customerID string) (body string, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		CustomerID: customerID,
	}

	endpoint := w.baseurl + "auth/getByCustomer"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth bycustomer: %w", err)
	}
	return
}

// AuthByDateRange fetches authentication info using a date range.
func (w Initializer) AuthByDateRange(page, limit, from, to string) (body string, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		From:  from,
		To:    to,
	}

	endpoint := w.baseurl + "auth/getByDate"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byDateRange: %w", err)
	}
	return
}

// AuthByBank fetches authentication info using the bank id.
func (w Initializer) AuthByBank(page, limit, bankID string) (body string, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		BankID: bankID,
	}

	endpoint := w.baseurl + "auth/getByBank"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byBank: %w", err)
	}
	return
}

// AuthByCustomerDate fetches authentication for a customer using a date range and customer id.
func (w Initializer) AuthByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		From:       from,
		To:         to,
		CustomerID: customerID,
	}

	endpoint := w.baseurl + "auth/getByDateCustomer"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byCustomerDate: %w", err)
	}
	return
}

/*
Balance Product, documentation can be found here https://docs.okra.ng/products/balance
*/

// RetrieveBalance retrieves Bank balance
func (w Initializer) RetrieveBalance() (body string, err error) {

	endpoint := w.baseurl + "products/balances"

	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving bank balance: %w", err)
	}
	return

}

// BalanceByID fetches balance info using the id of the balance.
func (w Initializer) BalanceByID(page, limit, i string) (body string, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		ID:    i,
	}

	endpoint := w.baseurl + "balance/getById"

	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching balance using id: %w", err)
	}
	return
}

// BalanceByOptions fetches balance info using the options metadata you provided when setting up the widget.
func (w Initializer) BalanceByOptions(page, limit, firstname, lastname string) (body string, err error) {

	pl := optionPayload{
		Page:  page,
		Limit: limit,
		Options: option{
			FirstName: firstname,
			LastName:  lastname,
		},
	}
	url := w.baseurl + "balance/byOptions"
	body, err = postRequest(pl, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance byoptions: %w", err)
	}
	return
}
