package translit_test

import (
	"fmt"

	"github.com/mxmCherry/translit"
)

func ExampleConverter() {
	tr := translit.New(map[string]string{
		"л":   "l",
		"Л":   "L",
		"лещ": "bream",
		"Лещ": "Bream",
	})

	// TODO: WOW, the first example failed if letter is last letter in string!
	fmt.Println(tr.Convert("Л - л"))                               // L - l
	fmt.Println(tr.Convert("Лещ - лещ"))                           // Bream - bream
	fmt.Println(tr.Convert("Не объявленные правиЛа игнорируются")) // Не объявlенные правиLа игнорируются

	// Output:
	// L - l
	// Bream - bream
	// Не объявlенные правиlа игнорируются
}
