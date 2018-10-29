package transformer_test

import (
	"github.com/mxmCherry/translit/internal/tree"
	"github.com/mxmCherry/translit/internal/tree/maptree"
	"golang.org/x/text/transform"

	. "github.com/mxmCherry/translit/internal/transformer"

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
	subject := New(mapping)

	BeforeEach(func() {
		subject.Reset()
	})

	DescribeTable(
		"Examples",

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

	It("should request more src if there is longer potential match", func() {
		var (
			nDst, nSrc int
			err        error
		)

		nDst, nSrc, err = subject.Transform(nil, []byte("A"), false)
		Expect(err).To(MatchError(transform.ErrShortSrc))
		Expect(nDst).To(BeZero())
		Expect(nSrc).To(BeZero())

		nDst, nSrc, err = subject.Transform(nil, []byte("AB"), false)
		Expect(err).To(MatchError(transform.ErrShortSrc))
		Expect(nDst).To(BeZero())
		Expect(nSrc).To(BeZero())
	})

	DescribeTable(
		"ErrShortSrc",

		func(input string) {
			nDst, nSrc, err := subject.Transform(nil, []byte("A"), false)
			Expect(err).To(MatchError(transform.ErrShortSrc))
			Expect(nDst).To(BeZero())
			Expect(nSrc).To(BeZero())
		},

		Entry("already has match, but also has longer potential match", "A"),
		Entry("no match", "AB"),
	)

	DescribeTable(
		"ErrShortDst",

		func(input string, atEOF bool) {
			nDst, nSrc, err := subject.Transform(nil, []byte(input), atEOF)
			Expect(err).To(MatchError(transform.ErrShortDst))
			Expect(nDst).To(BeZero())
			Expect(nSrc).To(BeZero())
		},

		Entry("!atEOF", "AC", false),
		Entry("atEOF, have match", "A", true),
		Entry("atEOF, no match", "D", true),
	)
})
