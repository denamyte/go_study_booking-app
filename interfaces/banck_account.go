package interfaces

import "fmt"

type BankAccount interface {
	GetBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Withdraw(amount int) error
}

func RunBankAccounts() {

	myAccounts := []BankAccount{
		NewWellsFargo(),
		NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdraw(270); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		balance := account.GetBalance()
		fmt.Printf("balance = %d\n", balance)
	}
}
