package main

import (
	"fmt"
	"os"

	okra "github.com/Uchencho/okraGo"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	ok := okra.NewOkra(token, "https://api.okra.ng/sandbox/v1/")

	body2, err2 := ok.RetrieveAuth()
	fmt.Println(err2, len(body2), body2[:50])
}
