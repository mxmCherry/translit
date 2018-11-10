package ruicao_test

import (
	"fmt"

	"github.com/mxmCherry/translit/ruicao"
	"golang.org/x/text/transform"
)

func ExampleToLatin() {
	ru := ruicao.ToLatin() // this is recommended to be a global variable in your own package

	// https://ru.wikipedia.org/wiki/Панграмма
	s, _, _ := transform.String(ru.Transformer(), "Съешь же ещё этих мягких французских булок да выпей чаю.")
	fmt.Println(s)

	// Output:
	// Sieesh zhe eshche etikh miagkikh frantsuzskikh bulok da vypei chaiu.
}
