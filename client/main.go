package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/okraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	okraClient := okra.NewOkra(token, "https://api.okra.ng/sandbox/v1/")

	body2, err2 := okraClient.RetrieveAuth()
	fmt.Println(err2, len(body2), body2[:50])
}
