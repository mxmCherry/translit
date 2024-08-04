package maptree_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMaptree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maptree Suite")
}
