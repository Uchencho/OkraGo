package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/okraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	// cID := "5e9d5dd3471ff50f735ad68a"
	okraClient := okra.New(token, "https://api.okra.ng/sandbox/v1/")

	// body2, err2 := okraClient.RetrieveAuth()
	// fmt.Println(err2, "\n\n", body2)

	// body, err := okraClient.TransactionByCustomer("1", "50", "5df1020beffe401cfaaa473f")
	// fmt.Println(err, body)

	// body3, err := okraClient.RetrieveBalance()
	// fmt.Println(err, "\n\n", body3)

	body, err := okraClient.RetrieveTransaction()
	fmt.Println(err, body.Data, body.Message)
}

// {"status":"success","message":"There are no incomes available for this client","data":{"pagination":{},"incomes":[]}}
