package maptree_test

import (
	"github.com/mxmCherry/translit/internal/tree"

	. "github.com/mxmCherry/translit/internal/tree/maptree"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	var _ tree.Node = (*Node)(nil)

	It("should lookup child", func() {
		childA := &Node{
			Val: []byte("{a}"),
		}
		childB := &Node{
			Val: []byte("{b}"),
		}
		subject := &Node{
			Children: map[byte]*Node{
				'a': childA,
				'b': childB,
			},
		}
		Expect(subject.LookupChild('a')).To(Equal(childA))
		Expect(subject.LookupChild('b')).To(Equal(childB))
		Expect(subject.LookupChild('x')).To(BeNil())
	})

	It("should return value", func() {
		Expect(
			(&Node{
				Val: []byte("{a}"),
			}).Value(),
		).To(Equal([]byte("{a}")))
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
					'a': {},
				},
			}).HasChildren(),
		).To(BeTrue())
	})

	It("should set value", func() {
		subject := &Node{}
		subject.SetValue([]byte("{a}"))
		Expect(subject.Val).To(Equal([]byte("{a}")))
	})

	It("should ensure child", func() {
		subject := &Node{}

		childA := subject.EnsureChild('a')
		childB := subject.EnsureChild('b')

		childA.SetValue([]byte("{a}"))
		childB.SetValue([]byte("{b}"))

		Expect(subject.EnsureChild('a')).To(Equal(childA))

		Expect(subject).To(Equal(&Node{
			Children: map[byte]*Node{
				'a': {Val: []byte("{a}")},
				'b': {Val: []byte("{b}")},
			},
		}))
	})
})
