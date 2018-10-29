package tree

// Node defines tree node.
type Node interface {
	Value() []byte
	LookupChild(byte) Node
	HasChildren() bool
}

// BuilderNode defines builder node.
type BuilderNode interface {
	SetValue([]byte)
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
