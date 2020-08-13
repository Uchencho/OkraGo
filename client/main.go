package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/OkraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")

	okraClient, err := okra.New(token, "https://api.okra.ng/sandbox/v1/")
	if err != nil {
		fmt.Println(err)
	}

	body2, err2 := okraClient.RetrieveAuth()
	fmt.Println(err2, body2.StatusCode)

}
