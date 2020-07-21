package main

import (
	"fmt"

	"github.com/Uchencho/okraGo/toy"
)

func main() {
	u := toy.Toy{
		Name:   "Uchenna Alozie",
		Weight: 50,
	}
	fmt.Println("Before calling New function:\t", u)
	toy.New(&u)

	fmt.Println("After calling New function:\t", u)
}
