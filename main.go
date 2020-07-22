package main

import (
	"fmt"

	"github.com/Uchencho/okraGo/toy"
)

func main() {
	u := toy.New("Uchenna", 50)
	u.UpdateOnHand(30)
	u.UpdateSold(10)

	fmt.Println(u.Name)
	fmt.Println(u.Weight)
	fmt.Println(u.Hand())
	fmt.Println(u.Sold())
}
