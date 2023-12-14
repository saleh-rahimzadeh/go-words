package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/saleh-rahimzadeh/go-words/core"
)

//__________________________________________________________________________________________________

// ValidationSource validate source string
func ValidationSource(source string) error {
	if source == Empty || strings.TrimSpace(source) == Empty {
		return core.ErrWordsEmpty
	}
	return nil
}

// ValidationDelimiters validate demilimters such as separator and comment
func ValidationDelimiters(separator string, comment string) error {
	if separator == comment {
		return core.ErrSameSeparatorAndComment
	}
	if !RegexSeparator.MatchString(separator) {
		return core.ErrSeparatorIsInvalid
	}
	if !RegexComments.MatchString(comment) {
		return core.ErrCommentIsInvalid
	}
	return nil
}

// ValidationName check name validation and return trimed name, return false if name is invalid
func ValidationName(name string) (string, bool) {
	name = strings.TrimSpace(name)
	if name == Empty {
		return Empty, false
	}
	if strings.Contains(name, NewLine) {
		return Empty, false
	}
	return name, true
}

// Extract search for a name in line and return value and true if found, else return empty string and false if not found
func Extract(line string, name string, separator string) (string, bool) {
	key, value, _ := strings.Cut(line, separator)
	if key == name {
		return value, true
	}
	return Empty, false
}

// Parse parse the line of words and return "key", "value" if has not error
func Parse(line string, separator string, comment string) (string, string, error) {
	var data = strings.TrimSpace(line)
	if data == Empty {
		return Empty, Empty, core.ErrLineEmpty
	}
	if strings.HasPrefix(data, comment) {
		return Empty, Empty, core.ErrLineComment
	}

	key, value, found := strings.Cut(data, separator)
	if !found {
		return Empty, Empty, fmt.Errorf("%w, at line '%s'", core.ErrSeparatorNotPresent, line)
	}

	key = strings.TrimSpace(key)
	if key == Empty {
		return Empty, Empty, fmt.Errorf("%w, at line '%s'", core.ErrNameNotPresent, line)
	}

	return key, strings.TrimSpace(value), nil
}

// NormalizeLine parse line and return prepared line
func NormalizeLine(line string, separator string, comment string) (string, error) {
	key, value, err := Parse(line, separator, comment)
	if err != nil {
		return Empty, err
	}
	return fmt.Sprintf("%s%s%s", key, separator, value), nil
}

// Normalization parse each line and return prepared source collection
func Normalization(source string, separator string, comment string) ([]string, error) {
	var collection = make([]string, 0, strings.Count(source, NewLine)+1)
	for _, line := range strings.Split(source, NewLine) {
		data, err := NormalizeLine(line, separator, comment)
		if err != nil {
			if errors.Is(err, core.ErrLineEmpty) || errors.Is(err, core.ErrLineComment) {
				continue
			}
			return nil, err
		}
		collection = append(collection, data)
	}
	return collection, nil
}

// Collection prepare source as a collection
func Collection(source []string, separator string) (map[string]string, error) {
	var collection = make(map[string]string, len(source)+1)
	for _, line := range source {
		key, value, _ := strings.Cut(line, separator)
		if _, found := collection[key]; found {
			return nil, fmt.Errorf("%w, name '%s'", core.ErrNameDuplicated, key)
		}
		collection[key] = value
	}
	return collection, nil
}

// CheckDuplication check for any duplicated names in source
func CheckDuplication(source []string, separator string) (bool, string) {
	for index, line := range source {
		compareKey, _, _ := strings.Cut(string(line), separator)
		for traverse := index + 1; traverse < len(source); traverse++ {
			comprandLine := source[traverse]
			comprandKey, _, _ := strings.Cut(string(comprandLine), separator)
			if compareKey == comprandKey {
				return true, compareKey
			}
		}
	}
	return false, Empty
}
