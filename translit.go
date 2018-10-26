// Package translit provides some basic utils for transliteration.
// See subpackages for language-specific converters.
package translit

import (
	"io"
	"strings"

	"github.com/mxmCherry/translit/internal/lookup"
)

// Converter represents transliteration converter.
type Converter interface {
	// Convert converts input string according to converter rules.
	// It is expected to get valid unicode strings.
	Convert(in string) (out string)
}

// New constructs case-sensitive converter from given transliteration map.
//
// TODO: OCD/naming woes: name it not New, but Build?..
func New(transliterations map[string]string) Converter {
	return converter{
		lookup: lookup.FromMap(transliterations),
	}
}

// ----------------------------------------------------------------------------

type converter struct {
	lookup lookup.Lookup
}

func (c converter) Convert(s string) string {
	in := strings.NewReader(s)

	var (
		out strings.Builder
		cur strings.Builder
	)
	var last struct {
		tr     *string
		curLen int
	}
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		_, _ = cur.WriteRune(r)

		tr, hasLongerMatch := c.lookup.Lookup(cur.String())

		if tr != nil {
			last.tr = tr
			last.curLen = cur.Len()
		}

		if hasLongerMatch && in.Len() > 0 {
			continue
		}

		if last.tr != nil {
			if rewind := cur.Len() - last.curLen; rewind > 0 {
				_, _ = in.Seek(-int64(rewind), io.SeekCurrent)
			}

			_, _ = out.WriteString(*last.tr)

			last.tr = nil
			last.curLen = 0
		} else {
			_, _ = out.WriteString(cur.String())
		}

		cur.Reset()
	}

	return out.String()
}
