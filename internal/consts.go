package internal

import "regexp"

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Pre-defined control characters
const (
	NewLine     string = "\n"
	NewLineByte byte   = '\n'
	Empty       string = ""
)

//┌ Regex
//└─────────────────────────────────────────────────────────────────────────────────────────────────

// Regex for separator and comment validation
var (
	RegexSeparator *regexp.Regexp = regexp.MustCompile(`^[=|:;,.?@]{1}$`)
	RegexComments  *regexp.Regexp = regexp.MustCompile(`^[#|:;,.?@]{1}$`)
)

//┌ Annotation
//└─────────────────────────────────────────────────────────────────────────────────────────────────

var RegexAnnotation *regexp.Regexp = regexp.MustCompile(`{{\s*(\w+)\s*}}`)

const AnnotationDelimiters string = "{} "
