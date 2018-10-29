package mapping_test

import (
	"github.com/mxmCherry/translit/internal/tree"
	"github.com/mxmCherry/translit/internal/tree/maptree"
	"golang.org/x/text/transform"

	. "github.com/mxmCherry/translit/internal/mapping"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transformer", func() {
	mapping := &maptree.Node{}
	tree.Populate(mapping, map[string]string{
		"A":   "{A}",
		"B":   "{B}",
		"C":   "{C}",
		"ABC": "{ABC}",
	})
	subject := NewTransformer(mapping)

	BeforeEach(func() {
		subject.Reset()
	})

	DescribeTable(
		"Common examples",

		func(input, expected string) {
			actual, n, err := transform.String(subject, input)
			Expect(err).NotTo(HaveOccurred())
			Expect(n).To(Equal(len(input)))
			Expect(actual).To(Equal(expected))
		},

		Entry(
			"single-char",
			"A",
			"{A}",
		),
		Entry(
			"all-same-char",
			"AAA",
			"{A}{A}{A}",
		),
		Entry(
			"generic",
			"ABAABCDAEA",
			"{A}{B}{A}{ABC}D{A}E{A}",
		),
	)

	PIt("should have 100% test coverage!")
})
