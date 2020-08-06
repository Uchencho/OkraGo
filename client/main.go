package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/OkraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	// cID := "5e9d5dd3471ff50f735ad68a"
	// const (
	// 	bankID = "5d6fe57a4099cc4b210bbeae"
	// 	bID    = "5d6fe57a4099cc4b210bbeb1"
	// )
	okraClient := okra.New(token, "https://api.okra.ng/sandbox/v1/")

	body2, err2 := okraClient.RetrieveAuth()
	fmt.Println(err2, body2.StatusCode)

	// body, err := okraClient.AuthByBank("1", "20", bankID)
	// fmt.Println(err, "\n\n", body.Message, body.Data.Auths)

	// body, code, err := okraClient.TransactionByID("1", "50", "5ed3d67b8723c11444c43283")
	// fmt.Println(err, body, code)

	// body, err := okraClient.TransactionBySpendingPattern("5defaec14ff9b30ae366483a")
	// fmt.Println(err, body)

	// body3, err := okraClient.BalanceByType("1", "20", "ledger_balance", "4000")
	// fmt.Println(err, "\n\n", body3)

	// body3, err := okraClient.PeriodicBalance("NGN", "5f231df19f28af16b6cf5994", "5f231e0c9e5c6e823adf606f")
	// fmt.Println(err, "\n\n", body3)

	// body3, err := okraClient.BalanceByOptions("1", "20", "James", "Galler")
	// fmt.Println(err, "\n\n", body3)

	body, code, err := okraClient.RetrieveIncome()
	fmt.Println("\n\n", err, body, code)

	// body2, err2 := okraClient.RetrieveTransaction()
	// fmt.Println(err2, body2.Data, body2.Message)

	// body, err := okraClient.IdentityByDateRange("1", "50", "2015-01-01", "2020-08-01")
	// fmt.Println(err, body)

	// body, err := okraClient.IdentityByDateRange("1", "20", "2015-01-01", "2020-08-01")
	// fmt.Println(err, body)

}

// {"status":"success","message":"There are no incomes available for this client","data":{"pagination":{},"incomes":[]}}
// Balance by type currently returns a 400 Bad request
// TransactionBySpendingPattern returns 400 bad request
// Retrieve Identities returning 504

/*
Struct Responses


Auth : 7, Done 6, remaining 1
Balance : 8, Done 6, remaining 2
Transaction : 11, Done 5, remaining 6
Identity and Income is currently unstable so returning string
*/
