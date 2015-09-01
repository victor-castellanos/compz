package snapshots_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSnapshots(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snapshots Suite")
}
