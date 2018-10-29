package maptree_test

import (
	"github.com/mxmCherry/translit/internal/experimental/tree"

	. "github.com/mxmCherry/translit/internal/experimental/tree/maptree"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	var _ tree.Node = (*Node)(nil)

	It("should insert", func() {
		subject := new(Node)

		subject.Insert([]byte("a"), []byte("{a}"))
		subject.Insert([]byte("ab"), []byte("{ab}"))
		subject.Insert([]byte("bc"), []byte("{bc}"))
		subject.Insert([]byte("cd"), []byte("{cd}"))
		subject.Insert([]byte("c"), []byte("{c}"))

		Expect(subject).To(Equal(
			&Node{
				Value: nil,
				Children: map[byte]*Node{
					'a': {
						Value: []byte("{a}"),
						Children: map[byte]*Node{
							'b': {Value: []byte("{ab}"), Children: nil},
						},
					},
					'b': {
						Value: nil,
						Children: map[byte]*Node{
							'c': {Value: []byte("{bc}"), Children: nil},
						},
					},
					'c': {
						Value: []byte("{c}"),
						Children: map[byte]*Node{
							'd': {Value: []byte("{cd}"), Children: nil},
						},
					},
				},
			},
		))
	})

	It("should tell if it has children", func() {
		Expect(
			(&Node{
				Children: nil,
			}).HasChildren(),
		).To(BeFalse())
		Expect(
			(&Node{
				Children: map[byte]*Node{},
			}).HasChildren(),
		).To(BeFalse())
		Expect(
			(&Node{
				Children: map[byte]*Node{
					'a': &Node{},
				},
			}).HasChildren(),
		).To(BeTrue())
	})

	It("should lookup child", func() {
		childA := &Node{
			Value: []byte("{a}"),
		}
		childB := &Node{
			Value: []byte("{b}"),
		}
		subject := &Node{
			Children: map[byte]*Node{
				'a': childA,
				'b': childB,
			},
		}
		Expect(subject.Child('a')).To(Equal(childA))
		Expect(subject.Child('b')).To(Equal(childB))
		Expect(subject.Child('x')).To(BeNil())
	})
})
