package wallet

import "fmt"

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{name: accountName}
}

func (account *Account) checkAccount(accountName string) error {
	if account.name != accountName {
		return fmt.Errorf("Account name %s does not match account name %s", account.name, accountName)
	}
	fmt.Printf("Account verified: %s\n", account.name)
	return nil
}
