package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/okraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	cID := "5ed3d67b8723c11444c43283"
	okraClient := okra.NewOkra(token, "https://api.okra.ng/sandbox/v1/")

	body2, err2 := okraClient.RetrieveAuth()
	fmt.Println(err2, len(body2), body2[:50])

	body, err := okraClient.TransactionByCustomer("1", "30", cID)
	fmt.Println(err, body)
}
