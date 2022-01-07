package test

import (
	"testing"

	ft "github.com/justinsowhat/french-tokenizer"
)

func TestTokenizerWithoutMergingProperNouns(t *testing.T) {

	tokenizer := ft.Tokenizer{}

	text := " «    Je m'appelle Jean-Pierre, et j'aime faire du foot justqu'au soir. » "
	mergeProperNouns := false
	actual := tokenizer.Tokenize(text, mergeProperNouns)
	if len(actual) != 22 {
		t.Fatalf("There should be 22 tokens!")
	}
}

func TestTokenizerWithMergingProperNouns(t *testing.T) {

	tokenizer := ft.Tokenizer{}

	text := "Jean-Pierre"
	mergeProperNouns := true
	actual := tokenizer.Tokenize(text, mergeProperNouns)
	if len(actual) != 1 {
		t.Fatalf("There should be 1 tokens!")
	}
}
func TestTokenizerURLHttps(t *testing.T) {

	tokenizer := ft.Tokenizer{}

	text := "https://www.github.com"
	mergeProperNouns := true
	actual := tokenizer.Tokenize(text, mergeProperNouns)
	if len(actual) != 1 {
		t.Fatalf("There should be 1 tokens!")
	}
}

func TestTokenizerURLHttp(t *testing.T) {

	tokenizer := ft.Tokenizer{}

	text := "http://www.github.com"
	mergeProperNouns := true
	actual := tokenizer.Tokenize(text, mergeProperNouns)
	if len(actual) != 1 {
		t.Fatalf("There should be 1 tokens!")
	}
}
