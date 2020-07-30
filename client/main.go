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

	body2, err2 := okraClient.RetrieveAuth()
	fmt.Println(err2, "\n\n", body2, "\n\n", body2.Status)

	// body, err := okraClient.TransactionByCustomer("1", "30", cID)
	// fmt.Println(err, body[:100])
}
