package dollar_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDollar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dollar Suite")
}
