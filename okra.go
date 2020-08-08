package okra

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Customer   string `json:"customer_id"`
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
func New(t, b string) (Client, error) {
	u := Client{
		token:   t,
		baseurl: b,
	}
	var err error
	if u.token == "" || u.baseurl == "" {
		err = errors.New("Token and Base url is needed to call this Function")
	} else if string(u.baseurl[len(u.baseurl)-1]) != "/" {
		u.baseurl = u.baseurl + "/"
	}
	return u, err
}

func postRequestByte(pl interface{}, url, token string) (body []byte, code int, err error) {

	reqBody, err := json.Marshal(pl)
	if err != nil {
		return body, 0, fmt.Errorf("error marshalling json: %w", err)
	}

	var bearer = "Bearer " + token
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return body, 0, fmt.Errorf("error making http call: %w", err)
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return body, 0, fmt.Errorf("error doing request: %w", err)
	}
	code = resp.StatusCode

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, code, fmt.Errorf("error reading body: %w", err)
	}

	return
}

// general unexported byid function since all 5 products have similar signature
func byID(page, limit, i, endpoint, token string) (body []byte, code int, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		ID:    i,
	}

	body, code, err = postRequestByte(pl, endpoint, token)
	if err != nil {
		return body, code, fmt.Errorf("error fetching product using id: %w", err)
	}
	return
}

// General byoptions function
func byOptions(page, limit, firstname, lastname, url, token string) (body string, code int, err error) {

	pl := optionPayload{
		Page:  page,
		Limit: limit,
		Options: option{
			FirstName: firstname,
			LastName:  lastname,
		},
	}
	bod, code, err := postRequestByte(pl, url, token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving product byoptions: %w", err)
	}
	return string(bod), code, err
}

func byCustomer(page, limit, customerID, endpoint, token string) (body []byte, code int, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		CustomerID: customerID,
	}

	body, code, err = postRequestByte(pl, endpoint, token)
	if err != nil {
		return body, code, fmt.Errorf("error retrieving product bycustomer: %w", err)
	}
	return
}

func byDateRange(page, limit, from, to, endpoint, token string) (body []byte, code int, err error) {

	pl := genPayload{
		Page:  page,
		Limit: limit,
		From:  from,
		To:    to,
	}
	body, code, err = postRequestByte(pl, endpoint, token)
	if err != nil {
		return body, code, fmt.Errorf("error retrieving product byDateRange: %w", err)
	}
	return
}

func byCustomerDate(page, limit, from, to, customerID, endpoint, token string) (body []byte, code int, err error) {

	pl := genPayload{
		Page:       page,
		Limit:      limit,
		From:       from,
		To:         to,
		CustomerID: customerID,
	}
	body, code, err = postRequestByte(pl, endpoint, token)
	if err != nil {
		return body, code, fmt.Errorf("error retrieving product byCustomerDate: %w", err)
	}
	return
}

// RetrieveAuth retrieves authentication of a user
func (w Client) RetrieveAuth() (body RetrieveAuthPayload, err error) {

	endpoint := w.baseurl + "products/auths"
	bod, code, err := postRequestByte(nil, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth token: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return

}

// AuthByID fetches authentication info using the id of the authentication record.
func (w Client) AuthByID(page, limit, ID string) (body AuthByIDPayload, err error) {

	endpoint := w.baseurl + "auth/getById"
	bod, code, err := byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error fetching auth using id: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// AuthByOptions fetches authentication info using the options metadata you provided when setting up the widget.
func (w Client) AuthByOptions(page, limit, firstname, lastname string) (body string, code int, err error) {

	url := w.baseurl + "auth/getByOptions"
	bod, code, err := byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving auth byoptions: %w", err)
	}
	return string(bod), code, err
}

// AuthByCustomer fetches authentication info using the customer id
func (w Client) AuthByCustomer(page, limit, customerID string) (body AuthByCustomerIDPayload, err error) {

	endpoint := w.baseurl + "auth/getByCustomer"
	bod, code, err := byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth bycustomer: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// AuthByDateRange fetches authentication info using a date range.
func (w Client) AuthByDateRange(page, limit, from, to string) (body AuthByDateRangePayload, err error) {

	endpoint := w.baseurl + "auth/getByDate"
	bod, code, err := byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth byDateRange: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// AuthByBank fetches authentication info using the bank id.
func (w Client) AuthByBank(page, limit, bankID string) (body AuthByBankPayload, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		BankID: bankID,
	}

	endpoint := w.baseurl + "auth/getByBank"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth byBank: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// AuthByCustomerDate fetches authentication for a customer using a date range and customer id.
func (w Client) AuthByCustomerDate(page, limit, from, to, customerID string) (body AuthByCustomerDateRangePayload, err error) {

	endpoint := w.baseurl + "auth/getByCustomerDate"
	bod, code, err := byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth byCustomerDate: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

/*
Balance Product, documentation can be found here https://docs.okra.ng/products/balance
*/

// RetrieveBalance retrieves Bank balance
func (w Client) RetrieveBalance() (body RetrieveBalancePayload, err error) {

	endpoint := w.baseurl + "products/balances"

	bod, code, err := postRequestByte(nil, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving bank balance: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// BalanceByID fetches balance info using the id of the balance.
func (w Client) BalanceByID(page, limit, ID string) (body BalanceByIDPayload, err error) {

	endpoint := w.baseurl + "balance/getById"
	bod, code, err := byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error fetching balance using id: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// BalanceByOptions fetches balance info using the options metadata you provided when setting up the widget.
func (w Client) BalanceByOptions(page, limit, firstname, lastname string) (body string, code int, err error) {

	url := w.baseurl + "balance/byOptions"
	bod, code, err := byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving balance byoptions: %w", err)
	}
	return string(bod), code, err
}

// BalanceByCustomer fetches balance info using the customer id
func (w Client) BalanceByCustomer(page, limit, customerID string) (body BalanceByCustomerIDPayload, err error) {

	endpoint := w.baseurl + "balance/getByCustomer"
	bod, code, err := byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving balance bycustomer: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// BalanceByAccount fetches balance info using the account id
func (w Client) BalanceByAccount(page, limit, AccountID string) (body BalanceByAccountPayload, err error) {

	pl := genPayload{
		Page:      page,
		Limit:     limit,
		AccountID: AccountID,
	}

	endpoint := w.baseurl + "balance/getByAccount"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving balance by accountID: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// BalanceByType fetches balance info using type of balance
func (w Client) BalanceByType(page, limit, theType, amount string) (body string, code int, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		Type:   theType,
		Amount: amount,
	}

	endpoint := w.baseurl + "balance/getByType"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving balance by type: %w", err)
	}
	return string(bod), code, err
}

// BalanceByCustomerDate fetches balance info of a customer using a date range and customer id.
func (w Client) BalanceByCustomerDate(page, limit, from, to, customerID string) (body BalanceByCustomerDateRangePayload, err error) {

	endpoint := w.baseurl + "balance/getByCustomerDate"
	bod, code, err := byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving balance byCustomerDate: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// PeriodicBalance fetches real-time BALANCE at anytime without heavy calculation of the transactions on each of an Record's accounts.
func (w Client) PeriodicBalance(currency, recordID, accountID string) (body PeriodicBalancePayload, err error) {

	pl := genPayload{
		Currency: currency,
		RecordID: recordID,
		Account:  accountID,
	}

	endpoint := w.baseurl + "products/balance/periodic"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving real time balance: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

/*
Transaction product, documentation can be found at https://docs.okra.ng/products/transactions
*/

// RetrieveTransaction retrieves transactions
func (w Client) RetrieveTransaction() (body RetrieveTransactionPayload, err error) {

	endpoint := w.baseurl + "products/transactions"
	bod, code, err := postRequestByte(nil, endpoint, w.token)
	if err != nil {
		return body, fmt.Errorf("error retrieving bank balance: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	body.StatusCode = code
	return
}

// TransactionByID fetches transaction info using the id of the transaction.
func (w Client) TransactionByID(page, limit, ID string) (body string, code int, err error) {

	endpoint := w.baseurl + "transaction/getById"
	bod, code, err := byID(page, limit, ID, endpoint, w.token)
	body = string(bod)
	if err != nil {
		return "Error", code, fmt.Errorf("error fetching Transaction using id: %w", err)
	}
	return
}

// TransactionByOptions fetches transaction info using the options metadata you provided when setting up the widget.
func (w Client) TransactionByOptions(page, limit, firstname, lastname string) (body string, code int, err error) {

	url := w.baseurl + "transaction/byOptions"
	bod, code, err := byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving transaction byoptions: %w", err)
	}
	return string(bod), code, err
}

// TransactionByCustomer fetches transaction info using the customer id
func (w Client) TransactionByCustomer(page, limit, customerID string) (body TransactionByCustomerIDPayload, err error) {

	endpoint := w.baseurl + "transaction/getByCustomer"
	bod, code, err := byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving balance bycustomer: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// TransactionByAccount fetches transaction info using the account id
func (w Client) TransactionByAccount(page, limit, AccountID string) (body string, code int, err error) {

	pl := genPayload{
		Page:      page,
		Limit:     limit,
		AccountID: AccountID,
	}

	endpoint := w.baseurl + "transaction/getByAccount"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving transaction by accountID: %w", err)
	}
	return string(bod), code, err
}

// TransactionByDateRange fetches transaction info using a date range.
func (w Client) TransactionByDateRange(page, limit, from, to string) (body TransactionByDateRangePayload, err error) {

	endpoint := w.baseurl + "transaction/getByDate"
	bod, code, err := byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving transaction byDateRange: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// TransactionByBank fetches transaction info using the bank id.
func (w Client) TransactionByBank(page, limit, bankID string) (body TransactionByBankIDPayload, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		BankID: bankID,
	}

	endpoint := w.baseurl + "transaction/getByBank"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving auth byBank: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// TransactionByType fetches transaction info using type of balance
func (w Client) TransactionByType(page, limit, theType, amount string) (body string, code int, err error) {

	pl := genPayload{
		Page:   page,
		Limit:  limit,
		Type:   theType,
		Amount: amount,
	}

	endpoint := w.baseurl + "transaction/getByType"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving Transaction by type: %w", err)
	}
	return string(bod), code, err
}

// TransactionBySpendingPattern fetches spending pattern of a customer.
func (w Client) TransactionBySpendingPattern(customerID string) (body string, code int, err error) {

	pl := genPayload{
		Customer: customerID,
	}

	endpoint := w.baseurl + "products/transactions/spending-pattern"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving spending pattern: %w", err)
	}
	return string(bod), code, err
}

// TransactionByCustomerDate fetches transaction info of a customer using a date range and customer id.
func (w Client) TransactionByCustomerDate(page, limit, from, to, customerID string) (body TransactionByCustomerDateRangePayload, err error) {

	endpoint := w.baseurl + "transaction/getByCustomerDate"
	pl := genPayload{
		Page:     page,
		Limit:    limit,
		From:     from,
		To:       to,
		Customer: customerID,
	}
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving Transaction byCustomerDate: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// PeriodicTransaction fetches real-time TRANSACTION at anytime on each of an Record's accounts.
func (w Client) PeriodicTransaction(currency, recordID, accountID string) (body string, code int, err error) {

	pl := genPayload{
		Currency: currency,
		RecordID: recordID,
		Account:  accountID,
	}

	endpoint := w.baseurl + "products/transactions/periodic"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving periodic transactions: %w", err)
	}
	return string(bod), code, err
}

/*
Identity product, documentation can be found at https://docs.okra.ng/products/identity
*/

// RetrieveIdentities retrieves various account holder information on file
func (w Client) RetrieveIdentities() (body RetrieveIdentitiesPayload, err error) {

	endpoint := w.baseurl + "products/identities"
	bod, code, err := postRequestByte(nil, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving account information: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// IdentityByID fetches various account holder information on file using the id
func (w Client) IdentityByID(page, limit, ID string) (body IdentityByCustomerPayload, err error) {

	endpoint := w.baseurl + "identity/getById"
	bod, code, err := byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error fetching identity using id: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// IdentityByOptions fetches identity info using the options metadata you provided when setting up the widget.
func (w Client) IdentityByOptions(page, limit, firstname, lastname string) (body string, code int, err error) {

	url := w.baseurl + "identity/byOptions"
	bod, code, err := byOptions(page, limit, firstname, lastname, url, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving identity byoptions: %w", err)
	}
	return string(bod), code, err
}

// IdentityByCustomer retrieve various account holder information on file using the customer id.
func (w Client) IdentityByCustomer(page, limit, customerID string) (body IdentityByCustomerPayload, err error) {

	endpoint := w.baseurl + "identity/getByCustomer"
	bod, code, err := byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving identity bycustomer: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// IdentityByDateRange fetches various account holder information on file using date range.
func (w Client) IdentityByDateRange(page, limit, from, to string) (body IdentityByCustomerDatePayload, err error) {

	endpoint := w.baseurl + "identity/getByDate"
	bod, code, err := byDateRange(page, limit, from, to, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving identity byDateRange: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// IdentityByCustomerDate fetches account holder information on file using date range and customer id.
func (w Client) IdentityByCustomerDate(page, limit, from, to, customerID string) (body IdentityByCustomerDatePayload, err error) {

	endpoint := w.baseurl + "identity/getByCustomerDate"
	bod, code, err := byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving identity byCustomerDate: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

/*
Income product, documentation can be found at https://docs.okra.ng/products/income
*/

// RetrieveIncome retrieves income record
func (w Client) RetrieveIncome() (body string, code int, err error) {

	endpoint := w.baseurl + "products/income/get"
	bod, code, err := postRequestByte(nil, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving income information: %w", err)
	}
	return string(bod), code, err
}

// IncomeByID retrieve information pertaining to a Record’s income using the id.
func (w Client) IncomeByID(page, limit, ID string) (body string, code int, err error) {

	endpoint := w.baseurl + "income/getById"
	bod, code, err := byID(page, limit, ID, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error fetching income using id: %w", err)
	}
	return string(bod), code, err
}

// IncomeByCustomer retrieve information pertaining to a Record’s income using the customer id.
func (w Client) IncomeByCustomer(page, limit, customerID string) (body string, code int, err error) {

	endpoint := w.baseurl + "income/getByCustomer"
	bod, code, err := byCustomer(page, limit, customerID, endpoint, w.token)
	if err != nil {
		return "Error", code, fmt.Errorf("error retrieving income bycustomer: %w", err)
	}
	return string(bod), code, err
}

// IncomeByCustomerDate retrieve information pertaining to a Record’s income using the customer id and date range.
func (w Client) IncomeByCustomerDate(page, limit, from, to, customerID string) (body IncomeByCustomerDatePayload, err error) {

	endpoint := w.baseurl + "income/getByCustomerDate"
	bod, code, err := byCustomerDate(page, limit, from, to, customerID, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error retrieving income byCustomerDate: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	body.StatusCode = code
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}

// ProcessIncome retrieves the income of particular customer using the customer's id.
func (w Client) ProcessIncome(customerID string) (body ProcessIncomePayload, err error) {

	pl := genPayload{
		CustomerID: customerID,
	}

	endpoint := w.baseurl + "products/income/process"
	bod, code, err := postRequestByte(pl, endpoint, w.token)
	if err != nil {
		body.StatusCode = code
		return body, fmt.Errorf("error processing income of a particular customer: %w", err)
	}
	err = json.Unmarshal(bod, &body)
	if err != nil {
		return body, fmt.Errorf("error Unmarshalling json: %w", err)
	}
	return
}
