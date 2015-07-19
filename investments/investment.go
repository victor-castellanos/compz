package investments

import (
	"github.com/victor-castellanos/compz/dollar"
)

type Investment struct {
	InterestRate float64
}

func (n Investment) GetInterestGenerated(amount dollar.Dollar) dollar.Dollar {
	value := amount.GetValue()
	return dollar.FromFloat64(value * (n.InterestRate / 100))
}

func (n Investment) CalculateCompoundInterest(amount dollar.Dollar, years int) dollar.Dollar {
	result := amount
	for i := 0; i < years; i++ {
		result = dollar.FromFloat64(n.GetInterestGenerated(result).GetValue() + result.GetValue())
	}
	return result
}
