package uknational

import (
	"github.com/mxmCherry/translit"

	"golang.org/x/text/transform"
)

// ToLatin returns Ukrainian romanization transform.Transformer,
// as defined in http://zakon.rada.gov.ua/laws/show/55-2010-п
func ToLatin() transform.Transformer {
	return toLatin
}

var toLatin transform.Transformer

var toLatinMap = map[string]string{
	"А": "A",
	"а": "a",

	"Б": "B",
	"б": "b",

	"В": "V",
	"в": "v",

	"Г": "H",
	"г": "h",

	"Ґ": "G",
	"ґ": "g",

	"Д": "D",
	"д": "d",

	"Е": "E",
	"е": "e",

	// на початку слова; в інших позиціях - див. `func init()`:
	"Є": "Ye",
	"є": "ye",

	"Ж": "Zh",
	"ж": "zh",

	"З": "Z",
	"з": "z",

	"И": "Y",
	"и": "y",

	"І": "I",
	"і": "i",

	// на початку слова; в інших позиціях - див. `func init()`:
	"Ї": "Yi",
	"ї": "yi",

	// на початку слова; в інших позиціях - див. `func init()`:
	"Й": "Y",
	"й": "y",

	"К": "K",
	"к": "k",

	"Л": "L",
	"л": "l",

	"М": "M",
	"м": "m",

	"Н": "N",
	"н": "n",

	"О": "O",
	"о": "o",

	"П": "P",
	"п": "p",

	"Р": "R",
	"р": "r",

	"С": "S",
	"с": "s",

	"Т": "T",
	"т": "t",

	"У": "U",
	"у": "u",

	"Ф": "F",
	"ф": "f",

	"Х": "Kh",
	"х": "kh",

	"Ц": "Ts",
	"ц": "ts",

	"Ч": "Ch",
	"ч": "ch",

	"Ш": "Sh",
	"ш": "sh",

	"Щ": "Shch",
	"щ": "shch",

	// на початку слова; в інших позиціях - див. `func init()`:
	"Ю": "Yu",
	"ю": "yu",

	// на початку слова; в інших позиціях - див. `func init()`:
	"Я": "Ya",
	"я": "ya",

	// 1. Буквосполучення "зг" відтворюється латиницею як "zgh"
	// (наприклад,  Згорани - Zghorany, Розгон - Rozghon) на
	// відміну від "zh" - відповідника української літери
	// "ж".
	"Зг": "Zgh", // title-case handling
	"ЗГ": "ZGH", // all-upper-case handling
	"зг": "zgh",

	// 2. М'який знак і апостроф латиницею не відтворюються.
	"ь": "",
	"Ь": "",
	"’": "", // https://uk.wikipedia.org/wiki/Апостроф
	"'": "", // https://uk.wikipedia.org/wiki/Апостроф
}

func init() {
	notBeginningOfWordTransliterations := map[string]string{
		"є": "ie",
		"ї": "i",
		"й": "i",
		"ю": "iu",
		"я": "ia",
	}

	// handling of special cases, when letter may be transliterated differently
	// depending on position in word (beginning / other positions);
	// brutal (memory-intensive, but not as much to bother), but does the trick:
	for _, a := range keys(toLatinMap) {
		for s, tr := range notBeginningOfWordTransliterations {
			toLatinMap[a+s] = toLatinMap[a] + tr
		}
	}

	toLatin = translit.New(toLatinMap)
}

func keys(m map[string]string) []string {
	kk := make([]string, 0, len(m))
	for k := range m {
		kk = append(kk, k)
	}
	return kk
}
