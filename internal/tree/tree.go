// Package tree defines a byte-to-byte-slice tree
// and implements some basic operations for it.
package tree

// Node defines tree node.
type Node interface {
	// Value returns node value.
	// It must return nil if there's no value held by current node (so empty `[]byte{}` if there's empty value).
	Value() []byte
	// LookupChild returns node child for given byte.
	// It must return nil if there's no child for given byte.
	LookupChild(byte) Node
	// HasChildren tells if node has at least one child.
	HasChildren() bool
}

// BuilderNode defines builder node.
type BuilderNode interface {
	// SetValue sets (overrides) node value.
	SetValue([]byte)
	// EnsureChild checks or pre-creates a child node for given byte and returns it.
	EnsureChild(byte) BuilderNode
}

// Insert builds key-value subtree for given root node.
func Insert(n BuilderNode, k, v []byte) {
	for _, b := range k {
		n = n.EnsureChild(b)
	}
	n.SetValue(v)
}

// Populate populates subtree from string-to-string map for given root node.
func Populate(n BuilderNode, m map[string]string) {
	for k, v := range m {
		Insert(n, []byte(k), []byte(v))
	}
}
