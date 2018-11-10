package translit_test

import (
	"fmt"

	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

func ExampleMap() {
	// pre-compile a transformer factory from rule map;
	// this is recommended to be a global variable in your own package:
	custom := translit.Map(
		map[string]string{
			"л":  "l",
			"Л":  "L",
			"ля": "lya",
			"Ля": "Lya",
		},
	)

	// get a "fresh" transformer (can be done for each transliteration instead of tr.Reset()):
	tr := custom.Transformer()

	var s string

	tr.Reset() // reset transformer state before usage - it is stateful and non-thread-safe
	s, _, _ = transform.String(tr, "Л - л")
	fmt.Println(s) // L - l

	tr.Reset() // reset transformer state before usage - it is stateful and non-thread-safe
	s, _, _ = transform.String(tr, "Ля-лЯ-ля")
	fmt.Println(s) // Lya-lЯ-lya - no rule for upper-case "Я", so it's not converted

	// Output:
	// L - l
	// Lya-lЯ-lya
}
