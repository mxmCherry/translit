package maptree

import "github.com/mxmCherry/translit/internal/experimental/tree"

// Node represents tree node with map lookup.
type Node struct {
	Value    []byte
	Children map[byte]*Node
}

// Insert inserts a key-value pair into tree, based off the current node.
func (n *Node) Insert(k []byte, v []byte) {
	for _, b := range k {
		c := n.Children[b]
		if c == nil {
			c = &Node{}
			if n.Children == nil {
				n.Children = map[byte]*Node{
					b: c,
				}
			} else {
				n.Children[b] = c
			}
		}
		n = c
	}
	n.Value = v
}

// Val implements tree.Node.
func (n *Node) Val() []byte {
	return n.Value
}

// Child implements tree.Node.
func (n *Node) Child(b byte) tree.Node {
	return n.Children[b]
}

// HasChildren implements tree.Node.
func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}
