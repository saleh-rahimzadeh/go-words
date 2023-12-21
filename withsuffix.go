package gowords

import (
	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// WithSuffix utilize Words interface with suffix to provide categorized words table and text resource
type WithSuffix struct { //EXTENDS: Words
	Words
	suffix string
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Get search for a name with suffix then return value if found, else return empty string
func (w WithSuffix) Get(name string) string {
	return w.Words.Get(name + w.suffix)
}

// Find search for a name with suffix then return value and `true` if found, else return empty string and `false`
func (w WithSuffix) Find(name string) (string, bool) {
	return w.Words.Find(name + w.suffix)
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewWordsRepository create a new instance of WithSuffix
func NewWithSuffix(words Words, suffix core.Suffix) (Words, error) {
	if words == nil {
		return nil, core.ErrWordsNil
	}

	strsuffix, ok := internal.ValidationSuffix(string(suffix))
	if !ok {
		return nil, core.ErrSuffixIsInvalid
	}

	return WithSuffix{
		Words:  words,
		suffix: strsuffix,
	}, nil
}
