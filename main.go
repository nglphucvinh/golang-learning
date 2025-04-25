package main

import (
	"awesomeProject/wallet"
	"fmt"
	"log"
)

func main() {
	walletFacade := wallet.NewWalletFacade("abc", 1234)
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		fmt.Println(err)
	}
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
