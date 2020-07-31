package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/okraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	// cID := "5e9d5dd3471ff50f735ad68a"
	const (
		bankID = "5d6fe57a4099cc4b210bbeae"
	)
	okraClient := okra.New(token, "https://api.okra.ng/sandbox/v1/")

	// body2, err2 := okraClient.RetrieveAuth()
	// fmt.Println(err2, "\n\n", body2.Data.Auths)

	body, err := okraClient.AuthByBank("1", "20", bankID)
	fmt.Println(err, "\n\n", body.Message, body.Data.Auths)

	// body, err := okraClient.TransactionByCustomer("1", "50", "5df1020beffe401cfaaa473f")
	// fmt.Println(err, body)

	// body3, err := okraClient.RetrieveBalance()
	// fmt.Println(err, "\n\n", body3)

	// body, err := okraClient.RetrieveTransaction()
	// fmt.Println(err, body.Data, body.Message)

	// body, err := okraClient.IdentityByDateRange("1", "50", "2015-01-01", "2020-08-01")
	// fmt.Println(err, body)

	// body, err := okraClient.IdentityByDateRange("1", "20", "2015-01-01", "2020-08-01")
	// fmt.Println(err, body)

}

// {"status":"success","message":"There are no incomes available for this client","data":{"pagination":{},"incomes":[]}}
