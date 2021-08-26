package main

import (
	"fmt"
	"roach/banking"
)

func main() {
	account := banking.Account{Owner: "roach", Balance: 10}
	fmt.Println(account)
}