package lookup_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLookup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lookup Suite")
}

// ----------------------------------------------------------------------------

func strptr(s string) *string {
	return &s
}
