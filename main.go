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
	code, body, _ := uc.RetrieveAuth()
	fmt.Println(code, len(body))
}
