package uknational_test

import (
	"fmt"

	"golang.org/x/text/transform"

	"github.com/mxmCherry/translit/uknational"
)

func ExampleToLatin() {
	uk := uknational.ToLatin()

	s, _, _ := transform.String(uk, "Український трансліт")
	fmt.Println(s) // Ukrainskyi translit

	// Output:
	// Ukrainskyi translit
}
