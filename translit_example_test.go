package translit_test

import (
	"fmt"

	"github.com/mxmCherry/translit/ruicao"
	"github.com/mxmCherry/translit/uknational"

	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

func ExampleCustom() {
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

func TODOExampleChain() {
	// TODO: this test fails - for some reason, it doesn't combine transliterators properly...
	//       Need to investigate later.

	var (
		uk = uknational.ToLatin()
		ru = ruicao.ToLatin()
	)
	var (
		ukRu = transform.Chain(uk.Transformer(), ru.Transformer()) // Ukrainian rules tried first
		ruUk = transform.Chain(ru.Transformer(), uk.Transformer()) // Russian rules tried first
	)

	var s string

	ukRu.Reset()
	s, _, _ = transform.String(ukRu, "гы")
	fmt.Println(s) // hy - "h" because Ukrainian rules tried first (it's "g" in Russian)

	ruUk.Reset()
	s, _, _ = transform.String(ruUk, "гі")
	fmt.Println(s) // gi - "g" because Russian rules tried first (it's "h" in Ukrainian)

	// Output:
	// hy
	// gi
}
