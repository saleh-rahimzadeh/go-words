package gowords

import (
	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type WithSuffix struct { //EXTENDS: Words
	Words
	suffix string
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func (w WithSuffix) Get(name string) string {
	return w.Words.Get(name + w.suffix)
}

func (w WithSuffix) Find(name string) (string, bool) {
	return w.Words.Find(name + w.suffix)
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

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
