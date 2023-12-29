package gowords

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

type WithAnnotation struct {
	Words
}

//┌ Public Methods
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func (w WithAnnotation) Get(name string, arguments map[string]string) string {
	value, _ := w.Find(name, arguments)
	return value
}

func (w WithAnnotation) Find(name string, arguments map[string]string) (string, bool) {
	value, found := w.Words.Find(name)
	if !found {
		return internal.Empty, false
	}
	if value == internal.Empty {
		return internal.Empty, true
	}
	return w.replacer(value, arguments), true
}

func (w WithAnnotation) GetIndexed(name string, arguments ...string) string {
	argumentMap := w.convertIndexesToMap(arguments)
	return w.Get(name, argumentMap)
}

func (w WithAnnotation) FindIndexed(name string, arguments ...string) (string, bool) {
	argumentMap := w.convertIndexesToMap(arguments)
	return w.Find(name, argumentMap)
}

func (w WithAnnotation) GetFormatted(name string, arguments ...interface{}) string {
	value, _ := w.FindFormatted(name, arguments...)
	return value
}

func (w WithAnnotation) FindFormatted(name string, arguments ...interface{}) (string, bool) {
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

func (w WithAnnotation) replacer(input string, arguments map[string]string) string {
	if arguments == nil {
		return input
	}
	return internal.RegexAnnotation.ReplaceAllStringFunc(input, func(token string) string {
		if value, ok := arguments[strings.Trim(token, internal.AnnotationDelimiters)]; ok {
			return value
		}
		return token
	})
}

func (w WithAnnotation) convertIndexesToMap(arguments []string) map[string]string {
	argumentMap := map[string]string{}
	for index, value := range arguments {
		argumentMap[strconv.Itoa(index+1)] = value
	}
	return argumentMap
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

func NewWithAnnotation(words Words) (WithAnnotation, error) {
	if words == nil {
		return WithAnnotation{}, core.ErrWordsNil
	}

	return WithAnnotation{
		Words: words,
	}, nil
}
