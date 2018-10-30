package uknational

import (
	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

// ToLatin returns Ukrainian romanization transform.Transformer,
// as defined in http://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF
//
// TODO: letter position handling - this is impossible to do 100% correctly with current implementation.
func ToLatin() transform.Transformer {
	return toLatin
}

var toLatin = translit.New(
	map[string]string{
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

		// poor man's position handling:
		"Є": "Ye", // на початку слова
		"є": "ie", // в інших позиціях

		"Ж": "Zh",
		"ж": "zh",

		"З": "Z",
		"з": "z",

		"И": "Y",
		"и": "y",

		"І": "I",
		"і": "i",

		// poor man's position handling:
		"Ї": "Yi", // на початку слова
		"ї": "i",  // в інших позиціях

		// poor man's position handling:
		"Й": "Y", // на початку слова
		"й": "i", // в інших позиціях

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

		// poor man's position handling:
		"Ю": "Yu", // на початку слова
		"ю": "iu", // в інших позиціях

		// poor man's position handling:
		"Я": "Ya", // на початку слова
		"я": "ia", // в інших позиціях

		// 1. Буквосполучення "зг" відтворюється латиницею як "zgh"
		// (наприклад,  Згорани - Zghorany, Розгон - Rozghon) на
		// відміну від "zh" - відповідника української літери
		// "ж".
		"Зг": "Zgh",
		"ЗГ": "ZGH", // all-upper-case handling
		"зг": "zgh",

		// 2. М'який знак і апостроф латиницею не відтворюються.
		"ь": "",
		"Ь": "",
		"’": "", // https://uk.wikipedia.org/wiki/%D0%90%D0%BF%D0%BE%D1%81%D1%82%D1%80%D0%BE%D1%84
		"'": "", // https://uk.wikipedia.org/wiki/%D0%90%D0%BF%D0%BE%D1%81%D1%82%D1%80%D0%BE%D1%84
	},
)
