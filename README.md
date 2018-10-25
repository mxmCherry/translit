# translit [![GoDoc](https://godoc.org/github.com/mxmCherry/translit?status.svg)](https://godoc.org/github.com/mxmCherry/translit) [![Build Status](https://travis-ci.org/mxmCherry/translit.svg?branch=master)](https://travis-ci.org/mxmCherry/translit) [![Go Report Card](https://goreportcard.com/badge/github.com/mxmCherry/translit)](https://goreportcard.com/report/github.com/mxmCherry/translit)
Go (Golang) utilities for (mostly Cyrillic) transliteration

# TODO

For v1:

- solve naming: probably, better not `uknational.Converter()`, but `uknational.Romanizer()` or so (just in case there will be latin-to-cyrillic converter at some point?)
- fill all the TODOs in README
- add Russian transliteration

Nice to have (though may be too hard / impossible to make an abstract implementation):

- letter position handling for (at least) `uknational`: it has custom transliteration rules for first letter, like `я`: `ya` (first) or `ia` (not first) - may become a part of `internal/lookup.Lookup.Lookup(substr, posInWord)`
- nicer case-handling (like "UPPER" and "Title" cased transliterations - requires lookahead, check [dchest/translit.go](https://github.com/dchest/translit.go))

## Features

- simple, but flexible configuration - `translit.New` accepts string-to-string transliteration rule map, so it should be really easy to implement/add custom transliterations
- properly handles multi-char rules, would work for latin-to-cyrillic transliterations like `z` -> `з` and `zh` -> `ж` (and not `zh` -> `зх`)
- expected decent performance - no regular expressions or anything like that (though that could actually be quite fast for transliteration use case)
- [The Unlicense](https://tldrlegal.com/license/unlicense): Anyone is free to copy, modify, publish, use, compile, sell, or distribute this software, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.

## Problems

- not flexible enough to differentiate transliteration rules for first/non-first word letter for now: for example, Ukrainain national transliteration standard differentiates `я` at the beginning of the word `ya` and `я` in the middle of the word `ia` (though this looks solveable, just a bit later; should be fine for common usage even with this "disadvantage")

## Usage

TODO: copy some example here.

## Guidelines

TODO: this is subject to change!

This package aims to provide default transliterations for some languages.

Subpackage names for these transliterations should be named as [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) language code + standard name, for example: `uknational`, where `uk` is the language code and `national` is a standard, defined by national government.

One package per language/standard recommendation is aiming reduce memory footprint: pay with memory only for what you actually use.

These packages should expose at least one `translit.Converter` (ideally - just one, again - for memory footprint). Functions are recommended to expose converters (to avoid accidental variable reassigning).

```go
package main

import "github.com/mxmCherry/translit/uknational"

var uk = uknational.Converter() // global one (package-local)

func main() {
	println(uk.Convert("Український трансліт"))
}
```

### Code style

- [editorconfig](https://editorconfig.org/) (recommended)
- [gofmt](https://blog.golang.org/go-fmt-your-code) (MUST)
- [goreportcard](https://goreportcard.com/report/github.com/mxmCherry/translit) (MUST)

## Motivation

At the moment of writing this lib, there were 3 analogs:

[gen1us2k/go-translit](github.com/gen1us2k/go-translit) at [46f1a0b])https://github.com/gen1us2k/go-translit/commit/46f1a0be552caadbdbc19cf9a6705c4402b7ab47):

- bad: only Latin -> Russian (depends on use case though)
- minor: unable to reuse for other languages
- very bad: anti-optimal at least  - looping through rules, compiling regexp and replacing entire string for each rule

[essentialkaos/translit](https://github.com/essentialkaos/translit) at [d62c0f9](https://github.com/essentialkaos/translit/commit/d62c0f98f9b32cda180f3e875d80a6afbaf34d9b):

- neutral: only Russian -> Latin
- good: really nice README, really nice code
- good: plenty of transliteration standards implemented
- good: has support for previous+current rune analysis for special case converting
- good: nice tests (probably)
- minor: too much "branding"
- bad (IMO): custom license (and legal stuff is always scary, at least for me)

[dchest/translit.go](https://github.com/dchest/translit.go) at [5528f11](https://github.com/dchest/translit.go/commit/5528f1177236f74b86bf5eecb7381bcda1074cba):

- neutral: only Russian -> Latin
- fatal: doesn't even compile with Go 1.11
- bad: last updated in 2011
- minor: single-rune conversion rules only (like, no way for special multi-char cases, though that may be relevant only for Ukrainian and/or Bulgarian, not sure)
- nice: attempt to handle cases like `ЩИ` -> `SCHI`, `Щи` -> `Schi`, though don't think it really matters / cannot be worked around with case-sensitive converter

Summary: [essentialkaos/translit](https://github.com/essentialkaos/translit) is non-flexible, but extremely nice for Russian -> Latin transliteration (plenty of standards implemented, handles special multi-char cases), but has custom license.
