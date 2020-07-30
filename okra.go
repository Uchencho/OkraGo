package okra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client struct
type Client struct {
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
	AccountID  string `json:"account"`
	Type       string `json:"type"`
	Amount     string `json:"amount"`
	Account    string `json:"account_id"`
	RecordID   string `json:"record_id"`
	Currency   string `json:"currency"`
}

// New returns a struct that can be used to call all methods. Panics if empty arguements are used.
func New(t, b string) Client {
	u := Client{
		token:   t,
		baseurl: b,
	}
	if u.token == "" || u.baseurl == "" {
		panic("Token and Base url is needed to call this Function")
	} else if string(u.baseurl[len(u.baseurl)-1]) != "/" {
		u.baseurl = u.baseurl + "/"
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
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

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
		body = "Status code returned was: " + strconv.Itoa(resp.StatusCode)
	} else {
		body = string(bod)
	}

	return
}

// general unexported byid function since all 5 products have similar signature
func byID(page, limit, i, endpoint, token string) (body string, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		ID:    i,
	}

	body, err = postRequest(pl, endpoint, token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching product using id: %w", err)
	}
	return
}

// General byoptions function
func byOptions(page, limit, firstname, lastname, url, token string) (body string, err error) {

	pl := optionPayload{
		Page:  page,
		Limit: limit,
		Options: option{
			FirstName: firstname,
			LastName:  lastname,
		},
	}
	body, err = postRequest(pl, url, token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving product byoptions: %w", err)
	}
	return
}

func byCustomer(page, limit, customerID, endpoint, token string) (body string, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		CustomerID: customerID,
	}

	body, err = postRequest(pl, endpoint, token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving product bycustomer: %w", err)
	}
	return
}

func byDateRange(page, limit, from, to, endpoint, token string) (body string, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		From:  from,
		To:    to,
	}
	body, err = postRequest(pl, endpoint, token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving product byDateRange: %w", err)
	}
	return
}

func byCustomerDate(page, limit, from, to, customerID, endpoint, token string) (body string, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		From:       from,
		To:         to,
		CustomerID: customerID,
	}
	body, err = postRequest(pl, endpoint, token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving product byCustomerDate: %w", err)
	}
	return
}

// RetrieveAuth retrieves authentication of a user
func (w Client) RetrieveAuth() (body string, err error) {

	endpoint := w.baseurl + "products/auths"
	// endpoint := getEndpointURI(w.baseurl, "products/auths")
	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth token: %w", err)
	}
	return

}

// AuthByID fetches authentication info using the id of the authentication record.
func (w Client) AuthByID(page, limit, ID string) (body string, err error) {

	endpoint := w.baseurl + "auth/getById"
	body, err = byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching auth using id: %w", err)
	}
	return
}

// AuthByOptions fetches authentication info using the options metadata you provided when setting up the widget.
func (w Client) AuthByOptions(page, limit, firstname, lastname string) (body string, err error) {

	url := w.baseurl + "auth/getByOptions"
	body, err = byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byoptions: %w", err)
	}
	return
}

// AuthByCustomer fetches authentication info using the customer id
func (w Client) AuthByCustomer(page, limit, customerID string) (body string, err error) {

	endpoint := w.baseurl + "auth/getByCustomer"
	body, err = byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth bycustomer: %w", err)
	}
	return
}

// AuthByDateRange fetches authentication info using a date range.
func (w Client) AuthByDateRange(page, limit, from, to string) (body string, err error) {

	endpoint := w.baseurl + "auth/getByDate"
	body, err = byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byDateRange: %w", err)
	}
	return
}

// AuthByBank fetches authentication info using the bank id.
func (w Client) AuthByBank(page, limit, bankID string) (body string, err error) {

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
func (w Client) AuthByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	endpoint := w.baseurl + "auth/getByCustomerDate"
	body, err = byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byCustomerDate: %w", err)
	}
	return
}

/*
Balance Product, documentation can be found here https://docs.okra.ng/products/balance
*/

// RetrieveBalance retrieves Bank balance
func (w Client) RetrieveBalance() (body string, err error) {

	endpoint := w.baseurl + "products/balances"

	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving bank balance: %w", err)
	}
	return

}

// BalanceByID fetches balance info using the id of the balance.
func (w Client) BalanceByID(page, limit, ID string) (body string, err error) {

	endpoint := w.baseurl + "balance/getById"
	body, err = byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching balance using id: %w", err)
	}
	return
}

// BalanceByOptions fetches balance info using the options metadata you provided when setting up the widget.
func (w Client) BalanceByOptions(page, limit, firstname, lastname string) (body string, err error) {

	url := w.baseurl + "balance/byOptions"
	body, err = byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance byoptions: %w", err)
	}
	return
}

// BalanceByCustomer fetches balance info using the customer id
func (w Client) BalanceByCustomer(page, limit, customerID string) (body string, err error) {

	endpoint := w.baseurl + "balance/getByCustomer"
	body, err = byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance bycustomer: %w", err)
	}
	return
}

// BalanceByAccount fetches balance info using the account id
func (w Client) BalanceByAccount(page, limit, AccountID string) (body string, err error) {

	pl := genPayload{
		Page:      page,
		Limit:     limit,
		AccountID: AccountID,
	}

	endpoint := w.baseurl + "balance/getByAccount"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance by accountID: %w", err)
	}
	return
}

// BalanceByType fetches balance info using type of balance
func (w Client) BalanceByType(page, limit, theType, amount string) (body string, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		Type:   theType,
		Amount: amount,
	}

	endpoint := w.baseurl + "balance/getByType"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance by type: %w", err)
	}
	return
}

// BalanceByCustomerDate fetches balance info of a customer using a date range and customer id.
func (w Client) BalanceByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	endpoint := w.baseurl + "balance/getByCustomerDate"
	body, err = byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance byCustomerDate: %w", err)
	}
	return
}

// PeriodicBalance fetches real-time BALANCE at anytime without heavy calculation of the transactions on each of an Record's accounts.
func (w Client) PeriodicBalance(currency, recordID, accountID string) (body string, err error) {

	pl := genPayload{
		Currency: currency,
		RecordID: recordID,
		Account:  accountID,
	}

	endpoint := w.baseurl + "products/balance/periodic"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving real time balance: %w", err)
	}
	return
}

/*
Transaction product, documentation can be found at https://docs.okra.ng/products/transactions
*/

// RetrieveTransaction retrieves transactions
func (w Client) RetrieveTransaction() (body string, err error) {

	endpoint := w.baseurl + "products/transactions"
	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving bank balance: %w", err)
	}
	return

}

// TransactionByID fetches transaction info using the id of the transaction.
func (w Client) TransactionByID(page, limit, ID string) (body string, err error) {

	endpoint := w.baseurl + "transaction/getById"
	body, err = byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching Transaction using id: %w", err)
	}
	return
}

// TransactionByOptions fetches transaction info using the options metadata you provided when setting up the widget.
func (w Client) TransactionByOptions(page, limit, firstname, lastname string) (body string, err error) {

	url := w.baseurl + "transaction/byOptions"
	body, err = byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving transaction byoptions: %w", err)
	}
	return
}

// TransactionByCustomer fetches transaction info using the customer id
func (w Client) TransactionByCustomer(page, limit, customerID string) (body string, err error) {

	endpoint := w.baseurl + "transaction/getByCustomer"
	body, err = byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving balance bycustomer: %w", err)
	}
	return
}

// TransactionByAccount fetches transaction info using the account id
func (w Client) TransactionByAccount(page, limit, AccountID string) (body string, err error) {

	pl := genPayload{
		Page:      page,
		Limit:     limit,
		AccountID: AccountID,
	}

	endpoint := w.baseurl + "transaction/getByAccount"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving transaction by accountID: %w", err)
	}
	return
}

// TransactionByDateRange fetches transaction info using a date range.
func (w Client) TransactionByDateRange(page, limit, from, to string) (body string, err error) {

	endpoint := w.baseurl + "transaction/getByDate"
	body, err = byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving transaction byDateRange: %w", err)
	}
	return
}

// TransactionByBank fetches transaction info using the bank id.
func (w Client) TransactionByBank(page, limit, bankID string) (body string, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		BankID: bankID,
	}

	endpoint := w.baseurl + "transaction/getByBank"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving auth byBank: %w", err)
	}
	return
}

// TransactionByType fetches transaction info using type of balance
func (w Client) TransactionByType(page, limit, theType, amount string) (body string, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		Type:   theType,
		Amount: amount,
	}

	endpoint := w.baseurl + "transaction/getByType"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving Transaction by type: %w", err)
	}
	return
}

// TransactionBySpendingPattern fetches spending pattern of a customer.
func (w Client) TransactionBySpendingPattern(customerID string) (body string, err error) {

	pl := genPayload{
		CustomerID: customerID,
	}

	endpoint := w.baseurl + "products/transactions/spending-pattern"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving spending pattern: %w", err)
	}
	return
}

// TransactionByCustomerDate fetches transaction info of a customer using a date range and customer id.
func (w Client) TransactionByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	endpoint := w.baseurl + "transaction/getByCustomerDate"
	body, err = byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving transaction byCustomerDate: %w", err)
	}
	return
}

// PeriodicTransaction fetches real-time TRANSACTION at anytime on each of an Record's accounts.
func (w Client) PeriodicTransaction(currency, recordID, accountID string) (body string, err error) {

	pl := genPayload{
		Currency: currency,
		RecordID: recordID,
		Account:  accountID,
	}

	endpoint := w.baseurl + "products/transactions/periodic"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving periodic transactions: %w", err)
	}
	return
}

/*
Identity product, documentation can be found at https://docs.okra.ng/products/identity
*/

// RetrieveIdentities retrieves various account holder information on file
func (w Client) RetrieveIdentities() (body string, err error) {

	endpoint := w.baseurl + "products/identities"
	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving account information: %w", err)
	}
	return
}

// IdentityByID fetches various account holder information on file using the id
func (w Client) IdentityByID(page, limit, ID string) (body string, err error) {

	endpoint := w.baseurl + "identity/getById"
	body, err = byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching identity using id: %w", err)
	}
	return
}

// IdentityByOptions fetches identity info using the options metadata you provided when setting up the widget.
func (w Client) IdentityByOptions(page, limit, firstname, lastname string) (body string, err error) {

	url := w.baseurl + "identity/byOptions"
	body, err = byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving identity byoptions: %w", err)
	}
	return
}

// IdentityByCustomer retrieve various account holder information on file using the customer id.
func (w Client) IdentityByCustomer(page, limit, customerID string) (body string, err error) {

	endpoint := w.baseurl + "identity/getByCustomer"
	body, err = byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving identity bycustomer: %w", err)
	}
	return
}

// IdentityByDateRange fetches various account holder information on file using date range.
func (w Client) IdentityByDateRange(page, limit, from, to string) (body string, err error) {

	endpoint := w.baseurl + "identity/getByDate"
	body, err = byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving identity byDateRange: %w", err)
	}
	return
}

// IdentityByCustomerDate fetches account holder information on file using date range and customer id.
func (w Client) IdentityByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	endpoint := w.baseurl + "identity/getByCustomerDate"
	body, err = byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving identity byCustomerDate: %w", err)
	}
	return
}

/*
Income product, documentation can be found at https://docs.okra.ng/products/income
*/

// RetrieveIncome retrieves income record
func (w Client) RetrieveIncome() (body string, err error) {

	endpoint := w.baseurl + "products/income/get"
	body, err = postRequest(nil, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving income information: %w", err)
	}
	return
}

// IncomeByID retrieve information pertaining to a Record’s income using the id.
func (w Client) IncomeByID(page, limit, ID string) (body string, err error) {

	endpoint := w.baseurl + "income/getById"
	body, err = byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error fetching income using id: %w", err)
	}
	return
}

// IncomeByCustomer retrieve information pertaining to a Record’s income using the customer id.
func (w Client) IncomeByCustomer(page, limit, customerID string) (body string, err error) {

	endpoint := w.baseurl + "income/getByCustomer"
	body, err = byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving income bycustomer: %w", err)
	}
	return
}

// IncomeByCustomerDate retrieve information pertaining to a Record’s income using the customer id and date range.
func (w Client) IncomeByCustomerDate(page, limit, from, to, customerID string) (body string, err error) {

	endpoint := w.baseurl + "income/getByCustomerDate"
	body, err = byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error retrieving income byCustomerDate: %w", err)
	}
	return
}

// ProcessIncome retrieve information pertaining to a Record’s income using the customer id and date range.
func (w Client) ProcessIncome(customerID string) (body string, err error) {

	pl := genPayload{
		CustomerID: customerID,
	}

	endpoint := w.baseurl + "products/income/process"
	body, err = postRequest(pl, endpoint, w.token)
	if err != nil {
		return "Error", fmt.Errorf("error processing income of a particular customer: %w", err)
	}
	return
}
