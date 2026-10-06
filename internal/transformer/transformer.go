// Package transformer implements a transliteration golang.org/x/text/transform.Transformer.
package transformer

import (
	"github.com/mxmCherry/translit/internal/tree"
	"golang.org/x/text/transform"
)

// Transformer implements transform.Transformer.
//
// It is stateful and not thread-safe.
type Transformer struct {
	rules tree.Node
}

// New creates new Transformer for given set of rules.
//
// Accepted rules are a tree.Node, which is expected to return nil Value() if there's no match.
func New(rules tree.Node) *Transformer {
	return &Transformer{
		rules: rules,
	}
}

// Reset implements transform.Transformer.
func (t *Transformer) Reset() {
}

// Transform implements transform.Transformer.
func (t *Transformer) Transform(dst, src []byte, atEOF bool) (int, int, error) {
	if len(src) == 0 {
		return 0, 0, nil
	}
	var match struct {
		value []byte
		pos   int
	}
	node := t.rules
	pos := 0

	for {
		if pos >= len(src) {
			break
		}

		node = node.LookupChild(src[pos])
		if node == nil {
			break
		}

		if v := node.Value(); v != nil {
			match.value = v
			match.pos = pos
		}

		pos++
	}

	if !atEOF {
		if node != nil {
			if node.HasChildren() {
				return 0, 0, transform.ErrShortSrc
			}
		}
	}

	if match.value != nil {
		if len(dst) < len(match.value) {
			return 0, 0, transform.ErrShortDst
		}
		return copy(dst, match.value), match.pos + 1, nil
	}

	if len(dst) < 1 {
		return 0, 0, transform.ErrShortDst
	}
	dst[0] = src[0]
	return 1, 1, nil
}
