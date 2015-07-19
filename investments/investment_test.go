package investments

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/victor-castellanos/compz/dollar"
)

var _ = Describe("Investment", func() {
	var (
		sut Investment
	)

	Describe("Get the compound interest generated over a period of time", func() {
		Context("With positive interest", func() {
			BeforeEach(func() {
				sut = Investment{
					InterestRate: 10,
				}
			})
			It("Should return the correct amount of interests", func() {
				Expect(sut.CalculateCompoundInterest(dollar.FromInt(10000), 2).GetValue()).To(Equal(12100.00))
			})
		})
	})

	Describe("Getting the value of a tic", func() {
		Context("With positive interest rate", func() {
			BeforeEach(func() {
				sut = Investment{
					InterestRate: 10,
				}
			})

			It("Should return the value of the interests generated", func() {
				Expect(sut.GetInterestGenerated(dollar.FromInt(10000))).To(Equal(dollar.FromInt(1000)))
			})
		})
		Context("With negative interest rate", func() {
			BeforeEach(func() {
				sut = Investment{
					InterestRate: -10,
				}
			})
			It("Should return a negative number generated", func() {
				Expect(sut.GetInterestGenerated(dollar.FromInt(10000))).To(Equal(dollar.FromInt(-1000)))
			})
		})
	})

})
