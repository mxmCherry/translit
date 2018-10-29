package uknational_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUknational(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uknational Suite")
}
