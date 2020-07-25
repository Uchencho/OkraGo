package main

import (
	// "github.com/Uchencho/okraGo/toy"
	"fmt"
	"os"

	"github.com/Uchencho/okraGo/auth"
)

func main() {

	token := os.Getenv("OKRA_TOKEN")
	uc := auth.New(token, "https://api.okra.ng/sandbox/v1/")
	body, err := uc.RetrieveAuth()
	fmt.Println(err, len(body), body[:50])
}
