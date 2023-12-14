package core

import (
	"errors"
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Pre-defined separator and comment character
const (
	Separator rune = '='
	Comment   rune = '#'
)

//┌ Errors
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// Errors for parsing source string and validations
var (
	ErrWords                   error = errors.New("error in words source")
	ErrWordsEmpty              error = errors.New("words source is empty")
	ErrNameNotPresent          error = errors.New("name not present in record")
	ErrNameDuplicated          error = errors.New("duplicated name found")
	ErrSameSeparatorAndComment error = errors.New("separator and comment are same character")
	ErrSeparatorNotPresent     error = errors.New("separator not present in record")
	ErrSeparatorIsInvalid      error = errors.New("separator character is invalid, the separator parameter must be one character delimiter of (=|:;,.?@)")
	ErrCommentIsInvalid        error = errors.New("comment character is invalid, the comment parameter must be one character delimiter of (#|:;,.?@)")
	ErrLineEmpty               error = errors.New("line is empty")
	ErrLineComment             error = errors.New("line is comment")
)
