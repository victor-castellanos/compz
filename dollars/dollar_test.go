package dollars

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dollar", func() {
	var (
		sut Dollar
	)

	Describe("Creating a Dollar object", func() {
		Context("From an integer", func() {
			It("Should return a dollar with the specified value", func() {
				Expect(FromInt(1000).GetValue()).To(Equal(1000.00))
			})
		})

		Context("From a float64", func() {

			BeforeEach(func() {
				sut = FromFloat64(float64(100.50))
			})
			It("Should return a dollar with the specified value", func() {
				Expect(sut.GetValue()).To(Equal(100.50))
			})
		})

		Context("From a string", func() {
			var (
				err error
			)
			Context("And it returns no errors", func() {
				BeforeEach(func() {
					sut, _ = FromString("1000")
				})

				It("Should return a dollar with the specified value", func() {
					Expect(sut.GetValue()).To(Equal(1000.00))
				})

			})

			Context("And it returns erros", func() {
				BeforeEach(func() {
					sut, err = FromString("Not valid dollar value")
				})

				It("Should return an error", func() {
					Ω(err).Should(HaveOccurred())
					Ω(sut.GetValue()).Should(Equal(float64(0)))
				})
			})

		})
	})

	Describe("Precision of a Dollar", func() {
		Context("When a value has more than 2 floating point precision", func() {
			BeforeEach(func() {
				sut = FromFloat64(1.356)
			})
			It("Should represent a floating point with two positions", func() {
				Expect(sut.GetValue()).To(Equal(1.36))
			})
		})
	})

	Describe("Get value from Dollar object", func() {
		BeforeEach(func() {
			sut, _ = FromString("1000")
		})

		It("Should return a rational number", func() {
			Expect(sut.GetValue()).To(Equal(1000.00))
		})
	})

	Describe("Sum dollar values", func() {
		It("Should return the result of adding two dollar values", func() {
			Expect(Sum(FromInt(1000), FromFloat64(1000))).To(Equal(FromFloat64(2000)))
		})
	})

})
