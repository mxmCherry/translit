// Package translit provides transliteration transformer.
// See subpackages for language-specific transformers.
package translit

import (
	"golang.org/x/text/transform"

	"github.com/mxmCherry/translit/internal/transformer"
	"github.com/mxmCherry/translit/internal/tree"
	"github.com/mxmCherry/translit/internal/tree/maptree"
)

// New creates case-sensitive text transform.Transformer from given mapping.
func New(mapping map[string]string) transform.Transformer {
	m := new(maptree.Node)
	tree.Populate(m, mapping)
	return transformer.New(m)
}
