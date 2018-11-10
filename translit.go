// Package translit provides transliteration transformer.
// See subpackages for language-specific transformers.
package translit

import (
	"golang.org/x/text/transform"

	"github.com/mxmCherry/translit/internal/transformer"
	"github.com/mxmCherry/translit/internal/tree"
	"github.com/mxmCherry/translit/internal/tree/maptree"
)

// Factory represents pre-compiled tree builder/factory.
type Factory interface {
	// Transformer returns/builds a "fresh" transliteration transformer.
	Transformer() transform.Transformer
}

type factory struct {
	mapping tree.Node
}

// Map pre-compiles string-to-string transliteration rule map into transformer factory.
func Map(mapping map[string]string) Factory {
	m := new(maptree.Node)
	tree.Populate(m, mapping)
	return &factory{
		mapping: m,
	}
}

func (f *factory) Transformer() transform.Transformer {
	return transformer.New(f.mapping)
}
