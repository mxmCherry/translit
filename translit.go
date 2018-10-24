// Package translit provides some basic utils for transliteration.
// See subpackages for language-specific converters.
package translit

import (
	"io"
	"strings"

	"github.com/mxmCherry/translit/internal/lookup"
)

// Lookup represents transliteration rule lookup.
// Case sensitivity (conversions like щ -> sch / Щ -> Sch / Щ -> SCH) may be solved as a part of lookup.
//
// TODO: add some case-handling wrappers for this.
type Lookup interface {
	Lookup(in string) (out *string, hasLongerMatch bool)
}

// Build builds case-sensitive transliteration rule lookup from string-to-string map.
func Build(transliterations map[string]string) Lookup {
	return lookup.FromMap(transliterations)
}

// ----------------------------------------------------------------------------

// Converter represents transliteration converter.
type Converter interface {
	// Convert converts input string according to converter rules.
	// It is expected to get valid unicode strings.
	Convert(in string) (out string)
}

// For constructs case-sensitive converter from given transliteration map.
func For(lookup Lookup) Converter {
	return converter{
		lookup: lookup,
	}
}

// ----------------------------------------------------------------------------

type converter struct {
	lookup Lookup
}

func (c converter) Convert(s string) string {
	in := strings.NewReader(s)

	var out strings.Builder

	var cur struct {
		buf   strings.Builder
		first rune
	}

	var last struct {
		tr     *string
		curLen int
	}
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if cur.buf.Len() == 0 {
			cur.first = r
		}
		_, _ = cur.buf.WriteRune(r)

		tr, hasLongerMatch := c.lookup.Lookup(cur.buf.String())

		if tr != nil {
			last.tr = tr
			last.curLen = cur.buf.Len()
		}

		if hasLongerMatch {
			continue
		}

		if last.tr != nil {
			if rewind := cur.buf.Len() - last.curLen; rewind > 0 {
				_, _ = in.Seek(-int64(rewind), io.SeekCurrent)
			}

			_, _ = out.WriteString(*last.tr)

			last.tr = nil
			last.curLen = 0
		} else {
			_, _ = out.WriteString(cur.buf.String())
		}

		cur.buf.Reset()
	}

	return out.String()
}
