package main

import (
	// "github.com/Uchencho/okraGo/toy"
	"fmt"

	"github.com/Uchencho/okraGo/auth"
)

func main() {

	// How the package will be consumed
	uc := auth.New("token", "https://baseurl/")
	uc.RetrieveAuth()
	fmt.Println(uc)
}
