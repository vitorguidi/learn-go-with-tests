package error_handling

import (
	"errors"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

var InsufficientFundsError = errors.New("insufficient funds for withdraw")

func (w *Wallet) Deposit(v BitCoin) {
	w.balance += v
}

func (w *Wallet) Withdraw(v BitCoin) error {
	if w.balance < v {
		return InsufficientFundsError
	}
	w.balance -= v
	return nil
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}
