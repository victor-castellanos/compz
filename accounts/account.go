package accounts

import (
	"github.com/victor-castellanos/compz/dollars"
)

type Account struct {
	Balance dollars.Dollar
	CapitalGain dollars.Dollar
}

func getInterestGenerated(amount dollars.Dollar, interestRate float64) dollars.Dollar {
	return dollars.FromFloat64(amount.GetValue() * (interestRate / 100))
}

type Investment struct{
	InterestRate float64
}

func (i Investment) Grow(account Account) Account {
	capitalGain := getInterestGenerated(account.Balance, i.InterestRate)
	return Account{
		CapitalGain: capitalGain,
		Balance: dollars.FromFloat64(capitalGain.GetValue() + account.Balance.GetValue()),
	}
}

