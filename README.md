# translit [![GoDoc](https://godoc.org/github.com/mxmCherry/translit?status.svg)](https://godoc.org/github.com/mxmCherry/translit) [![Build Status](https://travis-ci.org/mxmCherry/translit.svg?branch=master)](https://travis-ci.org/mxmCherry/translit) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit)
Go (Golang) utilities for (mostly Cyrillic) transliteration

# TODO

For v1:

- example for `uknational`, can be copied from [Приклади написання](http://zakon.rada.gov.ua/laws/show/55-2010-%D0%BF)
- settle the API (see Guidelines section) - probably, better have `uknational.Converter()` instead of `uknational.Lookup()` - this may be better to keep the stable API if doing some custom tweaks like letter position handling?
- fill all the TODOs in README
- add Russian transliteration

Nice to have (though may be too hard / impossible to make an abstract implementation):

- nicer case-handling (like "UPPER" and "Title" cased transliterations - requires lookahead, check [dchest/translit.go](https://github.com/dchest/translit.go))
- letter position handling (like `uknational`: has custom transliteration rules for first letter, like `я`: `ya` (first) or `ia` (not first))

## Features

TODO: describe features / advantages / shortcomings of main converter implementation. Also mention permissive license ('cause that's one of the main reasons for doing this lib).

## Usage

TODO: copy some example here.

## Guidelines

TODO: this is subject to change!

This package aims to provide default transliterations for some languages.

Subpackage names for these transliterations should be named as [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code + standard name, for example: `uknational`, where `uk` is the language code and `national` is a standard, defined by national government.

One package per language/standard recommendation is aiming reduce memory footprint: pay with memory only for what you actually use.

These packages should have at least `func Lookup() translit.Lookup` default getter, so usage would look like:

```go
package main

import (
	"github.com/mxmCherry/translit"
	"github.com/mxmCherry/translit/uknational"
)

var uk = translit.For(uknational.Lookup()) // global one

func main() {
	println(uk.Convert("Український трансліт"))
}
```

### Code style

- [editorconfig](https://editorconfig.org/) (recommended)
- [gofmt](https://blog.golang.org/go-fmt-your-code) (MUST)
- [goreportcard](https://goreportcard.com/report/github.com/mxmCherry/translit) (MUST)

## Motivation

TODO: describe existing analogs and explain why they are not perfect (don't forget license!).
