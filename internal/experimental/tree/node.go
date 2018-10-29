package tree

// Node represents a tree node.
type Node interface {
	Val() []byte
	Child(byte) Node
	HasChildren() bool
}
