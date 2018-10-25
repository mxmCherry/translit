# translit [![GoDoc](https://godoc.org/github.com/mxmCherry/translit?status.svg)](https://godoc.org/github.com/mxmCherry/translit) [![Build Status](https://travis-ci.org/mxmCherry/translit.svg?branch=master)](https://travis-ci.org/mxmCherry/translit) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit)
Go (Golang) utilities for (mostly Cyrillic) transliteration

## Features

TODO: describe features / advantages / shortcomings of main converter implementation.

## Usage

TODO: copy some example here.

## Guidelines

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

Code style:

- [editorconfig](https://editorconfig.org/) (recommended)
- [gofmt](https://blog.golang.org/go-fmt-your-code) (MUST)
- [goreportcard](https://goreportcard.com/report/github.com/mxmCherry/translit) (MUST)

## Motivation

TODO: describe existing analogs and explain why they are not perfect.
