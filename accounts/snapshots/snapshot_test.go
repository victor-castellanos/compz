package snapshots

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/victor-castellanos/compz/dollars"
)

var _ = Describe("Snapshot", func() {

	var (
		snapshot Snapshot
	)

	Describe("Creating a Snapshot", func() {

		BeforeEach(func() {
			snapshot = Snapshot{
				Year: 1,
				TotalGrowth: dollars.FromInt(20000),
				CurrentGrowth: dollars.FromInt(10000),
				TotalInvested: dollars.FromInt(30000),
				CurrentInvested: dollars.FromInt(25000),
			}
		})

		It("Should create the Snapshot with the correct year passed", func() {
			Expect(snapshot.Year).To(Equal(int32(1)))
		})

		It("Should create the Snapshot with the correct total growth", func() {
			Expect(snapshot.TotalGrowth).To(Equal(dollars.FromInt(20000)))
		})

		It("Should create the Snapshot with the correct year growth", func() {
			Expect(snapshot.CurrentGrowth).To(Equal(dollars.FromInt(10000)))
		})

		It("Should create the Snapshot with the correct invested amount", func() {
			Expect(snapshot.TotalInvested).To(Equal(dollars.FromInt(30000)))
		})

		It("Should create the Snapshot with the correct balance", func() {
			Expect(snapshot.CurrentInvested).To(Equal(dollars.FromInt(25000)))
		})
	})

	Describe("When parsing a Snapshot from JSON", func() {

	})

})
