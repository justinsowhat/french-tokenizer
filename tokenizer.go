package tokenizer

import (
	"fmt"
	"strings"

	regexp "github.com/dlclark/regexp2"
)

type Tokenizer interface {
	Tokenize(text string, mergeProperNouns bool) []string
}

type FrenchTokenizer struct {
}

func (t FrenchTokenizer) Tokenize(text string, mergeProperNouns bool) []string {
	// License: this regex pattern is copied from https://github.com/boudinfl/kea/blob/master/kea/kea.py#L63
	pattern := `(?:xumsi)(?:[lcdjmnts]|qu)[']|(http(s?):[^\s]+\.\w{2,3})|(\d+[.,]\d+)|([.-]+)|(\w+)|([^\w\s])`
	re := regexp.MustCompile(pattern, regexp.Multiline)

	text = strings.TrimSpace(text)

	var matches []string
	startIndex := 0
	m, _ := re.FindStringMatch(text)

	runeText := []rune(text)

	for m != nil {
		startIndex = m.Index
		matchedWord := string(runeText[startIndex : startIndex+m.Length])
		matches = append(matches, matchedWord)
		m, _ = re.FindNextMatch(m)
	}

	if mergeProperNouns {
		matches = t.mergeHypenatedPropernouns(matches)
	}

	return matches
}

// simple heuristic to merge hyphenated proper nouns
// it is case sensitive
func (t FrenchTokenizer) mergeHypenatedPropernouns(tokens []string) []string {

	var results []string

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "-" {
			if i >= 1 && i <= len(tokens)-1 {
				previousWord := tokens[i-1]
				nextWord := tokens[i+1]
				if t.shouldMerge(previousWord, nextWord) {
					results = results[:len(results)-1]
					results = append(results, fmt.Sprintf("%s-%s", previousWord, nextWord))
					i++
					continue
				}
			}
		}
		results = append(results, tokens[i])
	}

	return results
}

func (t FrenchTokenizer) shouldMerge(previousWord, nextWord string) bool {
	return t.isFirstLetterUpperCase(previousWord) && t.isFirstLetterUpperCase(nextWord)
}

func (t FrenchTokenizer) isFirstLetterUpperCase(word string) bool {
	re := regexp.MustCompile(`^[A-Z]\w.*`, 0)
	m, _ := re.FindStringMatch(word)
	return m != nil
}
