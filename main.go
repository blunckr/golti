package main

import (
	"fmt"
)

func main() {

	fmt.Println("vim-go")
	pubKey := MakeKeys()
	PrinkJwk(pubKey)
}
