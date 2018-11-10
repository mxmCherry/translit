package uknational_test

import (
	"fmt"

	"github.com/mxmCherry/translit/uknational"
	"golang.org/x/text/transform"
)

func ExampleToLatin() {
	uk := uknational.ToLatin() // this is recommended to be a global variable in your own package

	// https://uk.wikipedia.org/wiki/Панграма
	s, _, _ := transform.String(uk.Transformer(), "Десь чув, що той фраєр привіз їхньому царю грильяж та класну шубу з пір'я ґави.")
	fmt.Println(s)

	// Output:
	// Des chuv, shcho toi fraier pryviz yikhnomu tsariu hryliazh ta klasnu shubu z piria gavy.
}
