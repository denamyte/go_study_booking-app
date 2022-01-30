package interfaces

import "errors"

type WellsFargo struct {
	balance int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{}
}

func (wf *WellsFargo) GetBalance() int {
	return wf.balance
}

func (wf *WellsFargo) Deposit(amount int) {
	wf.balance += amount
}

func (wf *WellsFargo) Withdraw(amount int) error {
	newBalance := wf.balance - amount
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	wf.balance = newBalance
	return nil
}
