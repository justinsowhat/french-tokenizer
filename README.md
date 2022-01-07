[![Go](https://github.com/justinsowhat/french-tokenizer/actions/workflows/go.yml/badge.svg)](https://github.com/justinsowhat/french-tokenizer/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/justinsowhat/french-tokenizer.svg)](https://pkg.go.dev/github.com/justinsowhat/french-tokenizer)

# French Tokenizer

This is a basic rule-based tokenizer written in Go for French

## How to install
```
go get github.com/justinsowhat/french-tokenizer
```

## Usage

```
import (
    "fmt"

    ft "github.com/justinsowhat/french-tokenizer"
)

tokenizer := ft.Tokenizer{}

	text := " «    Je m'appelle Jean-Pierre, et j'aime faire du foot justqu'au soir. » "
	mergeProperNouns := false
	actual := tokenizer.Tokenize(text, mergeProperNouns)
	fmt.Println(actual)

```

When `mergeProperNouns` is set to `true`, proper nous like `Jean-Pierre` will not be broken into three separate tokens, like `Jean - Pierre`; instead, it will remain `Jean-Pierre`. Use it with caution, as it's written with basic heuristic.

## License
The MIT license is here [https://github.com/justinsowhat/french-tokenizer/blob/main/LICENSE].
