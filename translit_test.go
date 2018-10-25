package translit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/mxmCherry/translit"
)

var _ = Describe("Converter", func() {
	var subject Converter

	BeforeEach(func() {
		lookup := Build(map[string]string{
			"щ":  "sch",
			"я":  "ya",
			"ня": "nia",
		})
		subject = For(lookup)
	})

	DescribeTable(
		"Convert",
		func(in string, expected string) {
			Expect(subject.Convert(in)).To(Equal(expected))
		},

		Entry(
			"usual",
			"щя", "schya",
		),

		Entry(
			"multi-char source",
			"щнящ", "schniasch", // nia rule applied, longest match
		),

		Entry(
			"multi-char + included char",
			"яняя", "yaniaya",
		),

		Entry(
			"unknown chars left as-is",
			"шщъ", "шschъ",
		),
	)
})
