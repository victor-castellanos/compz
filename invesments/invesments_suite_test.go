package investments_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInvesments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Invesments Suite")
}
