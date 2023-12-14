package gowords

import (
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// WordsCollection provide words table and text resource with accepting string source and storing in map
type WordsCollection struct {
	collection map[string]string
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Get search for a name then return value if found, else return empty string
func (w WordsCollection) Get(name string) string {
	value, _ := w.Find(name)
	return value
}

// Find search for a name then return value and `true` if found, else return empty string and `false`
func (w WordsCollection) Find(name string) (string, bool) {
	name, ok := internal.ValidationName(name)
	if !ok {
		return internal.Empty, false
	}
	if value, found := w.collection[name]; found {
		return value, true
	}
	return internal.Empty, false
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewWordsCollection create a new instance of WordsCollection
func NewWordsCollection(source string, separator rune, comment rune) (WordsCollection, error) {
	var (
		separatorCharacter string = string(separator)
		commentCharacter   string = string(comment)
		err                error
	)

	err = internal.ValidationSource(source)
	if err != nil {
		return WordsCollection{}, err
	}

	err = internal.ValidationDelimiters(separatorCharacter, commentCharacter)
	if err != nil {
		return WordsCollection{}, err
	}

	repository, err := internal.Normalization(source, separatorCharacter, commentCharacter)
	if err != nil {
		return WordsCollection{}, err
	}

	collection, err := internal.Treasure(repository, separatorCharacter)
	if err != nil {
		return WordsCollection{}, err
	}

	return WordsCollection{
		collection: collection,
	}, nil
}
