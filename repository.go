package gowords

import (
	"fmt"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// WordsRepository provide words table and text resource with accepting string source and storing in array
type WordsRepository struct {
	repository []string
	separator  rune
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Get search for a name then return value if found, else return empty string
func (w WordsRepository) Get(name string) string {
	value, _ := w.Find(name)
	return value
}

// Find search for a name then return value and `true` if found, else return empty string and `false`
func (w WordsRepository) Find(name string) (string, bool) {
	name, ok := internal.ValidationName(name)
	if !ok {
		return internal.Empty, false
	}
	var separator = string(w.separator)
	for _, line := range w.repository {
		if value, found := internal.Extract(line, name, separator); found {
			return value, true
		}
	}
	return internal.Empty, false
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewWordsRepository create a new instance of WordsRepository
func NewWordsRepository(source string, separator rune, comment rune) (WordsRepository, error) {
	var (
		separatorCharacter string = string(separator)
		commentCharacter   string = string(comment)
		err                error
	)

	err = internal.ValidationSource(source)
	if err != nil {
		return WordsRepository{}, err
	}

	err = internal.ValidationDelimiters(separatorCharacter, commentCharacter)
	if err != nil {
		return WordsRepository{}, err
	}

	repository, err := internal.Normalization(source, separatorCharacter, commentCharacter)
	if err != nil {
		return WordsRepository{}, err
	}

	if duplicated, name := internal.CheckDuplication(repository, separatorCharacter); duplicated {
		return WordsRepository{}, fmt.Errorf("%w, name '%s'", core.ErrNameDuplicated, name)
	}

	return WordsRepository{
		repository: repository,
		separator:  separator,
	}, nil
}
