package gowords

import (
	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// GetBy a helper to search for a name by suffix and using Words object,
// then return value if found, else return empty string.
// Also return empty string if suffix is invalid.
func GetBy(w Words, name string, suffix core.Suffix) string {
	strsuffix, ok := internal.ValidationSuffix(string(suffix))
	if !ok {
		return internal.Empty
	}
	return w.Get(name + strsuffix)
}

// FindBy a helper to search for a name by suffix and using Words object,
// then return value and `true` if found, else return empty string and `false`.
// Also return empty string and `false` if suffix is invalid.
func FindBy(w Words, name string, suffix core.Suffix) (string, bool) {
	strsuffix, ok := internal.ValidationSuffix(string(suffix))
	if !ok {
		return internal.Empty, false
	}
	return w.Find(name + strsuffix)
}
