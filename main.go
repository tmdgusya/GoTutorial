package main

import (
	"fmt"
	"log"
	"roach/banking"
)

func main() {
	account := banking.NewAccount("roach")
	account.Deposit(1000)
	err := account.WithDraw(10000);
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}