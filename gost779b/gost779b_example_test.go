package gost779b_test

import (
	"fmt"

	"github.com/mxmCherry/translit"
	"github.com/mxmCherry/translit/gost779b"
)

func ExampleUK() {
	uk := translit.For(gost779b.UK())

	fmt.Println(uk.Convert("Український трансліт"))

	// Output:
	// Ukrayins`ky`j translit
}
