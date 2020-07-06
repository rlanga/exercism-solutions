package account

import (
	"sync"
)

type Account struct {
	balance int64
	open bool
	mux sync.Mutex
}

func Open(deposit int64) *Account {
	if deposit < 0 {
		return nil
	}
	return &Account{
		balance: deposit,
		open:    true,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open == false { return a.balance, false }
	a.open = false
	payout = a.balance
	a.balance = 0
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open == false { return a.balance, false }
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.open == false { return a.balance, false}

	if a.balance + amount < 0 { return a.balance, false}
	a.balance += amount
	return a.balance, true
}
