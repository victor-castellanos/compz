package accounts

import (
	"github.com/victor-castellanos/compz/dollars"
)

type Account struct {
	Balance     dollars.Dollar
	Invested    dollars.Dollar
	CapitalGain dollars.Dollar
}

func getInterestGenerated(amount dollars.Dollar, interestRate float64) dollars.Dollar {
	return dollars.FromFloat64(amount.GetValue() * (interestRate / 100))
}

type Investment struct {
	InterestRate float64
	Deposits     dollars.Dollar
}

func (self Investment) Grow(account Account) Account {
	initialBalance := dollars.Sum(account.Balance, self.Deposits)
	capitalGain := getInterestGenerated(initialBalance, self.InterestRate)
	return Account{
		CapitalGain: capitalGain,
		Invested:    dollars.Sum(self.Deposits, account.Invested),
		Balance:     dollars.Sum(capitalGain, initialBalance),
	}
}
