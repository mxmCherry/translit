package lookup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/mxmCherry/translit/internal/lookup"
)

var _ = Describe("Lookup", func() {
	var subject Lookup

	BeforeEach(func() {
		subject = FromMap(map[string]string{
			"abc": "ABC",
			"a":   "A",
			"b":   "B",
			"c":   "C",
		})
	})

	DescribeTable(
		"Lookup",

		func(key string, expectedValue *string, expectedHasLongerMatch bool) {
			actualValue, actualHasLongerMatch := subject.Lookup(key)
			Expect(actualValue).To(Equal(expectedValue))
			Expect(actualHasLongerMatch).To(Equal(expectedHasLongerMatch))
		},

		Entry(
			"no match",
			"xyz",
			nil, false,
		),

		Entry(
			"multi-char key",
			"abc",
			strptr("ABC"), false,
		),
		Entry(
			"first-char key",
			"a",
			strptr("A"), true, // has longer match (abc)
		),
		Entry(
			"middle-char key",
			"b",
			strptr("B"), false,
		),
		Entry(
			"last-char key",
			"c",
			strptr("C"), false,
		),
	)
})
