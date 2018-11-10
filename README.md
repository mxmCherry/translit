# translit [![GoDoc](https://godoc.org/github.com/mxmCherry/translit?status.svg)](https://godoc.org/github.com/mxmCherry/translit) [![Build Status](https://travis-ci.org/mxmCherry/translit.svg?branch=master)](https://travis-ci.org/mxmCherry/translit) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit) [![codecov](https://codecov.io/gh/mxmCherry/translit/branch/master/graph/badge.svg)](https://codecov.io/gh/mxmCherry/translit)

Go (Golang) utilities for (mostly Cyrillic) transliteration.

PROJECT STATE: pre-v1 grooming (mostly README).

This project aims to provide:

- easy-to-configure transliteration [golang.org/x/text/transform.Transformer](https://godoc.org/golang.org/x/text/transform#Transformer)
- language-specific transliterations (only Ukrainian/National and Russian/ICAO for now, PRs/issues are welcome)

This project is intended to be used with [golang.org/x/text/transform](https://godoc.org/golang.org/x/text/transform) - well-thought-out/convenient base for streaming text transforming.

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
- [codecov](https://codecov.io/gh/mxmCherry/translit) (recommended)

## Motivation

TODO: review/groom this a bit

At the moment of writing this lib, there were [3 analogs](https://golanglibs.com/search?q=translit):

[gen1us2k/go-translit](github.com/gen1us2k/go-translit) at [46f1a0b](https://github.com/gen1us2k/go-translit/commit/46f1a0be552caadbdbc19cf9a6705c4402b7ab47):

- bad: only Latin → Russian (depends on use case though)
- minor: unable to reuse for other languages
- very bad: anti-optimal at least  - looping through rules, compiling regexp and replacing entire string for each rule

[essentialkaos/translit](https://github.com/essentialkaos/translit) at [d62c0f9](https://github.com/essentialkaos/translit/commit/d62c0f98f9b32cda180f3e875d80a6afbaf34d9b):

- neutral: only Russian → Latin
- good: really nice README, really nice code
- good: plenty of transliteration standards implemented
- good: has support for previous+current rune analysis for special case converting
- good: nice tests (probably)
- minor: too much "branding"
- bad (IMO): custom license (and legal stuff is always scary, at least for me)

[dchest/translit.go](https://github.com/dchest/translit.go) at [5528f11](https://github.com/dchest/translit.go/commit/5528f1177236f74b86bf5eecb7381bcda1074cba):

- neutral: only Russian → Latin
- fatal: doesn't even compile with Go 1.11
- bad: last updated in 2011
- minor: single-rune conversion rules only (like, no way for special multi-char cases, though that may be relevant only for Ukrainian and/or Bulgarian, not sure)
- nice: attempt to handle cases like `ЩИ` → `SCHI`, `Щи` → `Schi`, though don't think it really matters / cannot be worked around with case-sensitive converter

Summary: [essentialkaos/translit](https://github.com/essentialkaos/translit) is non-flexible, but extremely nice for Russian → Latin transliteration (plenty of standards implemented, handles special multi-char cases), but has custom license; others are dead/inefficient.
