package gowords

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// WordsFile provide words table and text resource with accepting file pointer and storing a pointer to the file
type WordsFile struct {
	file      *os.File
	separator rune
	comment   rune
	fault     error
	mutex     *sync.Mutex
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Get search for a name then return value if found, else return empty string
func (w WordsFile) Get(name string) string {
	value, _ := w.Find(name)
	return value
}

// Find search for a name then return value and `true` if found, else return empty string and `false`.
// It is safe for concurrent use by multiple goroutines.
func (w WordsFile) Find(name string) (value string, found bool) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.FindUnsafe(name)

}

// FindUnsafe search for a name then return value and `true` if found, else return empty string and `false`.
// It is unsafe for concurrent use by multiple goroutines.
func (w *WordsFile) FindUnsafe(name string) (value string, found bool) {
	name, ok := internal.ValidationName(name)
	if !ok {
		return internal.Empty, false
	}

	defer func() {
		if rec := recover(); rec != nil {
			value = internal.Empty
			found = false
			if err, ok := rec.(error); ok {
				w.fault = err
			} else {
				w.fault = core.ErrWords
			}
		}
	}()

	_, err := w.file.Seek(0, 0)
	if err != nil {
		w.fault = err
		return internal.Empty, false
	}

	var (
		separatorCharacter string = string(w.separator)
		commentCharacter   string = string(w.comment)
		line               string
	)

	var scanner = bufio.NewScanner(w.file)
	for scanner.Scan() {
		line = scanner.Text()
		key, value, err := internal.Parse(line, separatorCharacter, commentCharacter)
		if err != nil {
			if errors.Is(err, core.ErrLineEmpty) || errors.Is(err, core.ErrLineComment) {
				continue
			}
			break
		}
		if key == name {
			return value, true
		}
	}
	return internal.Empty, false
}

// CheckError check errors in file.
// Also check for duplication of names.
func (w *WordsFile) CheckError() error {
	_, err := w.file.Seek(0, 0)
	if err != nil {
		return err
	}

	var (
		separatorCharacter string              = string(w.separator)
		commentCharacter   string              = string(w.comment)
		names              map[string]struct{} = make(map[string]struct{})
		scanner            *bufio.Scanner      = bufio.NewScanner(w.file)
		line               string
	)

	for scanner.Scan() {
		line = scanner.Text()
		key, _, err := internal.Parse(line, separatorCharacter, commentCharacter)
		if err != nil {
			if errors.Is(err, core.ErrLineEmpty) || errors.Is(err, core.ErrLineComment) {
				continue
			}
			return err
		}
		if _, found := names[key]; found {
			return fmt.Errorf("%w, name '%s'", core.ErrNameDuplicated, key)
		}
		names[key] = struct{}{}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

// Err get the error occurred in "Find" method
func (w *WordsFile) Err() error {
	return w.fault
}

//──────────────────────────────────────────────────────────────────────────────────────────────────

// NewWordsFile create a new instance of WordsFile
func NewWordsFile(file *os.File, separator rune, comment rune) (w WordsFile, err error) {
	var (
		separatorCharacter string = string(separator)
		commentCharacter   string = string(comment)
	)

	err = internal.ValidationDelimiters(separatorCharacter, commentCharacter)
	if err != nil {
		return WordsFile{}, err
	}

	return WordsFile{
		file:      file,
		separator: separator,
		comment:   comment,
		mutex:     &sync.Mutex{},
	}, nil
}
