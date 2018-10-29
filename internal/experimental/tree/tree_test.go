package tree_test

import (
	. "github.com/mxmCherry/translit/internal/experimental/tree"
	"github.com/mxmCherry/translit/internal/experimental/tree/maptree"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Insert", func() {
	It("should build subtree", func() {
		root := &maptree.Node{}

		Insert(root, []byte("a"), []byte("{a}"))
		Insert(root, []byte("ab"), []byte("{ab}"))
		Insert(root, []byte("bc"), []byte("{bc}"))
		Insert(root, []byte("cd"), []byte("{cd}"))
		Insert(root, []byte("c"), []byte("{c}"))
		Insert(root, []byte(""), []byte("{}"))

		Expect(root).To(Equal(
			&maptree.Node{
				Val: []byte("{}"),
				Children: map[byte]*maptree.Node{
					'a': {
						Val: []byte("{a}"),
						Children: map[byte]*maptree.Node{
							'b': {Val: []byte("{ab}"), Children: nil},
						},
					},
					'b': {
						Val: nil,
						Children: map[byte]*maptree.Node{
							'c': {Val: []byte("{bc}"), Children: nil},
						},
					},
					'c': {
						Val: []byte("{c}"),
						Children: map[byte]*maptree.Node{
							'd': {Val: []byte("{cd}"), Children: nil},
						},
					},
				},
			},
		))
	})
})

var _ = Describe("Populate", func() {
	It("should build subtree", func() {
		root := &maptree.Node{}

		Populate(root, map[string]string{
			"a":  "{a}",
			"ab": "{ab}",
			"bc": "{bc}",
			"cd": "{cd}",
			"c":  "{c}",
			"":   "{}",
		})

		Expect(root).To(Equal(
			&maptree.Node{
				Val: []byte("{}"),
				Children: map[byte]*maptree.Node{
					'a': {
						Val: []byte("{a}"),
						Children: map[byte]*maptree.Node{
							'b': {Val: []byte("{ab}"), Children: nil},
						},
					},
					'b': {
						Val: nil,
						Children: map[byte]*maptree.Node{
							'c': {Val: []byte("{bc}"), Children: nil},
						},
					},
					'c': {
						Val: []byte("{c}"),
						Children: map[byte]*maptree.Node{
							'd': {Val: []byte("{cd}"), Children: nil},
						},
					},
				},
			},
		))
	})
})
