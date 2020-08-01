package okra_test

import (
	"os"
	"testing"

	okra "github.com/Uchencho/okraGo"
)

const (
	baseurl = "https://api.okra.ng/sandbox/v1"
	succeed = "\u2713"
	failed  = "\u2717"
)

var token = os.Getenv("OKRA_TOKEN")

func TestInitializeOkra(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Okra New did not panic as a result of sending an empty string")
		}
	}()
	_ = okra.New("", baseurl)

}

func TestRetrieveAuth(t *testing.T) {
	okraClient := okra.New(token, baseurl)
	body, err := okraClient.RetrieveAuth()
	if err != nil || body.StatusCode != 200 {
		t.Fatalf("\t%s\tTest RetrieveAuth:\tGot Error: %v, and statuscode is : %v, Expected %v", failed, err, body.StatusCode, nil)
	}
	t.Logf("\t%s\tTest RetrieveAuth:\tShould have returned no errors.", succeed)
}

func TestTransactionByCustomerID(t *testing.T) {
	okraClient := okra.New(token, baseurl)
	cID := "5e9d5dd3471ff50f735ad68a"
	body, err := okraClient.TransactionByCustomer("1", "30", cID)
	if err != nil || body.StatusCode != 200 {
		t.Fatalf("\t%s\tTest TransactionByCustomer:\tGot Error: %v, and statuscode : %v, Expected %v", failed, err, body.StatusCode, nil)
	}
	t.Logf("\t%s\tTest TransactionByCustomer:\tShould have returned no errors.", succeed)
}
