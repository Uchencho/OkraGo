package okra_test

import (
	"os"
	"testing"

	okra "github.com/Uchencho/OkraGo"
)

// type mockingCalls interface {
// 	// https://blog.learngoprogramming.com/how-to-mock-in-your-go-golang-tests-b9eee7d7c266
// 	RetrieveAuth() okra.RetrieveAuthPayload
// }

const (
	baseurl = "https://api.okra.ng/sandbox/v1"

	succeed = "\u2713"
	failed  = "\u2717"
)

var token = os.Getenv("OKRA_TOKEN")

func TestInitializeOkra(t *testing.T) {

	const errorMessage = "Token and Base url is needed to call this Function"
	_, err := okra.New("", baseurl)
	if err == nil {
		t.Fatalf("\t%s\tTest NewClient:\tGot Error: %v, Expected %v", failed, nil, errorMessage)
	}
	t.Logf("\t%s\tTest NewClient:\tShould have returned error %v", succeed, err)

}

func TestRetrieveAuth(t *testing.T) {
	okraClient, err2 := okra.New(token, baseurl)
	if err2 != nil {
		body, err := okraClient.RetrieveAuth()
		if err != nil || body.StatusCode != 200 {
			t.Fatalf("\t%s\tTest RetrieveAuth:\tGot Error: %v, and statuscode is : %v, Expected %v", failed, err, body.StatusCode, nil)
		}
		t.Logf("\t%s\tTest RetrieveAuth:\tShould have returned no errors.", succeed)
	}
}

func TestTransactionByCustomerID(t *testing.T) {
	okraClient, _ := okra.New(token, baseurl)
	cID := "5e9d5dd3471ff50f735ad68a"
	body, err := okraClient.TransactionByCustomer("1", "30", cID)
	if err != nil || body.StatusCode != 200 {
		t.Fatalf("\t%s\tTest TransactionByCustomer:\tGot Error: %v, and statuscode : %v, Expected %v", failed, err, body.StatusCode, nil)
	}
	t.Logf("\t%s\tTest TransactionByCustomer:\tShould have returned no errors.", succeed)
}
