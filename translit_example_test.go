package translit_test

import (
	"fmt"

	"github.com/mxmCherry/translit"
	"golang.org/x/text/transform"
)

func ExampleNew() {
	tr := translit.New(map[string]string{
		"л":  "l",
		"Л":  "L",
		"ля": "lya",
		"Ля": "Lya",
	})

	var s string

	s, _, _ = transform.String(tr, "Л - л")
	fmt.Println(s) // L - l

	s, _, _ = transform.String(tr, "Ля-лЯ-ля")
	fmt.Println(s) // Lya-lЯ-lya

	// Output:
	// L - l
	// Lya-lЯ-lya
}
