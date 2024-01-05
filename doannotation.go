package gowords

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// DoAnnotation utilize Words interface to provide words table and text resource and format value according to an annotation or a format specifier
type DoAnnotation struct {
	Words
}

//┌ Public Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// GetNamed search for a name then return value if found, else return empty string.
// Format value with named annotations.
func (w DoAnnotation) GetNamed(name string, arguments map[string]interface{}) string {
	value, _ := w.FindNamed(name, arguments)
	return value
}

// FindNamed search for a name then return value and `true` if found, else return empty string and `false`.
// Format value with named annotations.
func (w DoAnnotation) FindNamed(name string, arguments map[string]interface{}) (string, bool) {
	value, found := w.Words.Find(name)
	if !found {
		return internal.Empty, false
	}
	if value == internal.Empty {
		return internal.Empty, true
	}
	return w.replacer(value, arguments), true
}

// GetIndexed search for a name then return value if found, else return empty string.
// Format value with indexed annotations.
func (w DoAnnotation) GetIndexed(name string, arguments ...interface{}) string {
	argumentMap := w.convertIndexesToMap(arguments)
	value, _ := w.FindNamed(name, argumentMap)
	return value
}

// FindIndexed search for a name then return value and `true` if found, else return empty string and `false`.
// Format value with indexed annotations.
func (w DoAnnotation) FindIndexed(name string, arguments ...interface{}) (string, bool) {
	argumentMap := w.convertIndexesToMap(arguments)
	return w.FindNamed(name, argumentMap)
}

// GetNamed search for a name then return value if found, else return empty string.
// Format value with formatted verbs according to "https://pkg.go.dev/fmt#hdr-Printing".
func (w DoAnnotation) GetFormatted(name string, arguments ...interface{}) string {
	value, _ := w.FindFormatted(name, arguments...)
	return value
}

// FindFormatted search for a name then return value and `true` if found, else return empty string and `false`.
// Format value with formatted verbs according to "https://pkg.go.dev/fmt#hdr-Printing".
func (w DoAnnotation) FindFormatted(name string, arguments ...interface{}) (string, bool) {
	value, found := w.Words.Find(name)
	if !found {
		return internal.Empty, false
	}
	if value == internal.Empty {
		return internal.Empty, true
	}
	return fmt.Sprintf(value, arguments...), true
}

//┌ Private Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// replacer finds annotation tokens in input and replace with genuine value
func (w DoAnnotation) replacer(input string, arguments map[string]interface{}) string {
	if arguments == nil {
		return input
	}
	return internal.RegexAnnotation.ReplaceAllStringFunc(input, func(token string) string {
		if value, ok := arguments[strings.Trim(token, internal.AnnotationDelimiters)]; ok {
			return fmt.Sprint(value)
		}
		return token
	})
}

// convertIndexesToMap converts indexed annotations to map
func (w DoAnnotation) convertIndexesToMap(arguments []interface{}) map[string]interface{} {
	argumentMap := map[string]interface{}{}
	for index, value := range arguments {
		argumentMap[strconv.Itoa(index+1)] = value
	}
	return argumentMap
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewDoAnnotation create a new instance of NewDoAnnotation
func NewDoAnnotation(words Words) (DoAnnotation, error) {
	if words == nil {
		return DoAnnotation{}, core.ErrWordsNil
	}

	return DoAnnotation{
		Words: words,
	}, nil
}
