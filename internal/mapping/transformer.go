package mapping

import (
	"github.com/mxmCherry/translit/internal/tree"
	"golang.org/x/text/transform"
)

// Transformer is a mapping transform.Transformer.
type Transformer struct {
	mapping tree.Node
	state   struct {
		node tree.Node
		pos  int
	}
}

// NewTransformer creates new Transformer for given mapping.
func NewTransformer(mapping tree.Node) *Transformer {
	return &Transformer{
		mapping: mapping,
	}
}

// Reset implements transform.Transformer.
func (t *Transformer) Reset() {
	t.state.node = nil
}

// Transform implements transform.Transformer.
func (t *Transformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	var match struct {
		value []byte
		pos   int
	}

	if t.state.node == nil {
		t.state.node = t.mapping
		t.state.pos = 0
	}

	for {
		if t.state.pos >= len(src) {
			break
		}

		t.state.node = t.state.node.LookupChild(src[t.state.pos])
		if t.state.node == nil {
			break
		}

		if v := t.state.node.Value(); v != nil {
			match.value = v
			match.pos = t.state.pos
		}

		t.state.pos++
	}

	if !atEOF {
		if t.state.node != nil {
			if t.state.node.HasChildren() {
				return 0, 0, transform.ErrShortSrc
			}
		}
	}

	if match.value != nil {
		if len(dst) < len(match.value) {
			return 0, 0, transform.ErrShortDst
		}
		t.state.node = nil
		return copy(dst, match.value), match.pos + 1, nil
	}

	if len(dst) < 1 {
		return 0, 0, transform.ErrShortDst
	}
	t.state.node = nil
	dst[0] = src[0]
	return 1, 1, nil
}
