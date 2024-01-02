# translit

[![Go Reference](https://pkg.go.dev/badge/github.com/mxmCherry/translit.svg)](https://pkg.go.dev/github.com/mxmCherry/translit)
[![Go Test](https://github.com/mxmCherry/translit/actions/workflows/go.yml/badge.svg)](https://github.com/mxmCherry/translit/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit)

Go (Golang) utilities for (mostly Cyrillic) transliteration.

This project aims to provide:

- easy-to-configure transliteration [golang.org/x/text/transform.Transformer](https://pkg.go.dev/golang.org/x/text/transform#Transformer)
- language-specific transliterations (only Ukrainian/National and Russian/ICAO for now, PRs/issues are welcome)

This project is intended to be used with [golang.org/x/text/transform](https://pkg.go.dev/golang.org/x/text/transform) - well-thought-out/convenient base for streaming text transforming.

If you a) don't need to build a custom transliterator b) are fine with custom license - take a look at [essentialkaos/translit](https://github.com/essentialkaos/translit) - it is fast and has plenty of standards implemented for Cyrillic.

## Features

- easy to build custom transliterations with `map[string]string` rules - `translit.Map(map[string]string{...}).Transformer()`
- tries longest match first - correct multi-character transliterations like `z` → `з` and `zh` → `ж` (and not `zh` → `зх`)
- expected decent performance - implemented using tree lookup under the hood (no regexps, no brute-force multi-replacement; transformer itself should not generate any garbage on it's own)
- [The Unlicense](https://tldrlegal.com/license/unlicense): Anyone is free to copy, modify, publish, use, compile, sell, or distribute this software, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.

It is intentionally kept as simple as possible, but that comes at cost of a higher memory usage in some special cases:

- rules like "first word character" / "non-first word character" result in multiple tree paths (roughly, each "alphabet letter" + "non-first letter" permutations - applies mostly to `uknational` implementation)
- upper/lower case - each upper/lower character conversion should be a custom rule, like `Б` → `B` and `б` → `b`

## Usage

```shell
go get -u github.com/mxmCherry/translit
```

### Custom rules

```go
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
```

### Language-specific transliterator

```go
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
```

## Guidelines

This package aims to provide default transliterations for some languages.

Subpackage names for these transliterations should be made of [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code and the standard name, for example: `uknational`, where `uk` is the language code and `national` is a standard (defined by national government).

One subpackage per language/standard approach is to reduce memory footprint: pay (with memory) only for what you actually use.

These subpackages should expose at least one `translit.Factory` constructor, ideally - two (two-way transformers, like `ToLatin()`/`FromLatin()`).

### Code style

- [editorconfig](https://editorconfig.org/) (recommended)
- [gofmt](https://blog.golang.org/go-fmt-your-code) (MUST)
- [goreportcard](https://goreportcard.com/report/github.com/mxmCherry/translit) (MUST)
