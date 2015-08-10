package accounts

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/victor-castellanos/compz/dollars"
)

var _ = Describe("Investment", func() {
	var (
		sut     Investment
		account Account
	)

	Describe("Grow", func() {
		BeforeEach(func() {
			sut = Investment{
				InterestRate: 10,
			}
			account = sut.Grow(Account{
				Balance: dollars.FromInt(10000),
			})
		})

		It("should return an account with the growth balance", func() {
			Expect(account.Balance).To(Equal(dollars.FromInt(11000)))
			Expect(account.CapitalGain).To(Equal(dollars.FromInt(1000)))
		})

		Context("with periodic addition", func() {
			BeforeEach(func() {
				sut.Deposits = dollars.FromInt(10000)
			})

			Context("after multiple tics", func() {
				BeforeEach(func() {
					account = sut.Grow(sut.Grow(Account{}))
				})

				It("should add a deposit on every tic", func() {
					Expect(account.Balance).To(Equal(dollars.FromInt(23100)))
					Expect(account.Invested).To(Equal(dollars.FromInt(20000)))
				})
			})
		})
	})

})
