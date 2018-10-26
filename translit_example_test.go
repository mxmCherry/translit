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

	fmt.Println(tr.Convert("Л - л"))                                // L - l
	fmt.Println(tr.Convert("Лещ - лещ - лЕщ"))                      // Bream - bream - lЕщ
	fmt.Println(tr.Convert("Остальной текст не трансЛитерируется")) // Остаlьной текст не трансLитерируется

	// Output:
	// L - l
	// Bream - bream - lЕщ
	// Остаlьной текст не трансLитерируется
}
