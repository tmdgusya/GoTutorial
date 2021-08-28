package main

import (
	"fmt"
	"roach/banking"
)

func main() {
	account := banking.NewAccount("roach")
	fmt.Println(account)
}