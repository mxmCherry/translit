package ruicao

import (
	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

// ToLatin returns Russian romanization transform.Transformer, according to ICAO standard:
//
// Doc 9303: Machine Readable Travel Documents (https://www.icao.int/publications/pages/publication.aspx?docnum=9303)
//
// Part 3: Specifications Common to all MRTDs (https://www.icao.int/publications/Documents/9303_p3_cons_en.pdf)
//
// Section 6. TRANSLITERATIONS RECOMMENDED FOR USE BY STATES
//
// Table A. Transliteration of Multinational Latin-based Characters
func ToLatin() transform.Transformer {
	return toLatin
}

var toLatin = translit.New(
	// https://en.wikipedia.org/wiki/Romanization_of_Russian#Transliteration_table
	map[string]string{
		"А": "A",
		"а": "a",

		"Б": "B",
		"б": "b",

		"В": "V",
		"в": "v",

		"Г": "G",
		"г": "g",

		"Д": "D",
		"д": "d",

		"Е": "E",
		"е": "e",

		"Ё": "E",
		"ё": "e",

		"Ж": "Zh",
		"ж": "zh",

		"З": "Z",
		"з": "z",

		"И": "I",
		"и": "i",

		"Й": "I",
		"й": "i",

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

		"Ъ": "Ie",
		"ъ": "ie",

		"Ы": "Y",
		"ы": "y",

		"Ь": "",
		"ь": "",

		"Э": "E",
		"э": "e",

		"Ю": "Iu",
		"ю": "iu",

		"Я": "Ia",
		"я": "ia",
	},
)
