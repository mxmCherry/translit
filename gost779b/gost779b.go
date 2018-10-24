// Package gost779b provides GOST 7.79 System B transliteration rules/lookup
// https://en.wikipedia.org/wiki/ISO_9#GOST_7.79_System_B
//
// These are intended to be used with root translit package:
//
//   import (
//     "github.com/mxmCherry/translit"
//     "github.com/mxmCherry/translit/gost779b"
//   )
//
//   uk := translit.For(gost779b.UK())
//
package gost779b

import (
	"strings"

	"github.com/mxmCherry/translit"
)

// BG returns (builds - so better get this only in `init()`) Bulgarian-flavored GOST 7.79 System B transliteration rule lookup.
func BG() translit.Lookup {
	m := cpy(base)
	m["Щ"] = "Sht"
	m["Ъ"] = "A`"
	populateLower(m)
	return translit.Build(m)
}

// BY returns (builds - so better get this only in `init()`) Belarusian-flavored GOST 7.79 System B transliteration rule lookup.
func BY() translit.Lookup {
	return defaultLookup()
}

// MK returns (builds - so better get this only in `init()`) Macedonian-flavored GOST 7.79 System B transliteration rule lookup.
func MK() translit.Lookup {
	return defaultLookup()
}

// RU returns (builds - so better get this only in `init()`) Russian-flavored GOST 7.79 System B transliteration rule lookup.
func RU() translit.Lookup {
	return defaultLookup()
}

// UK returns (builds - so better get this only in `init()`) Ukrainian-flavored GOST 7.79 System B transliteration rule lookup.
func UK() translit.Lookup {
	m := cpy(base)
	m["И"] = "Y`"
	populateLower(m)
	return translit.Build(m)
}

// ----------------------------------------------------------------------------

// TODO: probably, better to re-populate from https://ru.wikipedia.org/wiki/ISO_9#%D0%A2%D0%B0%D0%B1%D0%BB._2._%D0%A2%D1%80%D0%B0%D0%BD%D1%81%D0%BB%D0%B8%D1%82%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D1%8F_%D0%BF%D0%BE_%D1%81%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D0%B5_%D0%91 table

func defaultLookup() translit.Lookup {
	m := cpy(base)
	populateLower(m)
	return translit.Build(m)
}

func cpy(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func populateLower(m map[string]string) {
	for k, v := range m {
		m[strings.ToLower(k)] = strings.ToLower(v)
	}
}

var base = map[string]string{
	"А": "A",
	"Б": "B",
	"В": "V",
	"Г": "G",
	"Ѓ": "G`",
	"Ґ": "G`",
	"Д": "D",
	"Е": "E",
	"Ё": "Yo",
	"Є": "Ye",
	"Ж": "Zh",
	"З": "Z",
	"S": "Z`",
	"И": "I", // not in Belarusian, y` for Ukrainian (handled in UK())
	"Й": "J",
	"Ј": "J",
	"І": "I", // i` only before vowels for Old Russian and Old Bulgarian // TODO: add these exceptions at some point
	"Ї": "Yi",
	"К": "K",
	"Ќ": "K`",
	"Л": "L",
	"Љ": "L`",
	"М": "M",
	"Н": "N",
	"Њ": "N`",
	"О": "O",
	"П": "P",
	"Р": "R",
	"С": "S",
	"Т": "T",
	"У": "U",
	"Ў": "U`",
	"Ф": "F",
	"Х": "X",
	"Ц": "Cz", // c before i, e, y, j // TODO: add exceptions at some point
	"Ч": "Ch",
	"Џ": "Dh",
	"Ш": "Sh",
	"Щ": "Shh", // shh for Russian and Ukrainian, sht for Bulgarian (handled in BG())
	"Ъ": "``",
	"Ы": "Y`",
	"Ь": "`",
	"Э": "E`",
	"Ю": "Yu",
	"Я": "Ya",
	"’": "'",
	"Ѣ": "Ye",
	"Ѳ": "Fh",
	"Ѵ": "Yh",
	"Ѫ": "O`",
}
