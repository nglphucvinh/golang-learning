package wallet

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	return &WalletFacade{
		account:      newAccount(accountId),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}

func (w *WalletFacade) AddMoneyToWallet(accountId string, code, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.validateAccountAndCode(accountId, code)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountId string, code, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.validateAccountAndCode(accountId, code)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "debit", amount)
	return nil
}

// Validate account and security code
func (w *WalletFacade) validateAccountAndCode(accountId string, code int) error {
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}
	return nil
}
