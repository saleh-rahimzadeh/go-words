package gowords_test

import (
	"os"
	"path"
	"testing"

	. "github.com/saleh-rahimzadeh/go-words"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//┌ Test
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func TestNewWithSuffix(t *testing.T) {
	// Arrange
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	sourceFile, err := os.Open(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	defer sourceFile.Close()
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wFile, err := NewWordsFile(sourceFile, core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	const EN core.Suffix = "_EN"
	// Act
	_, errRepository := NewWithSuffix(wRepository, EN)
	_, errCollection := NewWithSuffix(wCollection, EN)
	_, errFile := NewWithSuffix(wFile, EN)
	// Assert
	if errRepository != nil {
		t.Errorf("NewWithSuffix(WordsRepository) error = %v", err)
		return
	}
	if errCollection != nil {
		t.Errorf("NewWithSuffix(WordsCollection) error = %v", err)
		return
	}
	if errFile != nil {
		t.Errorf("NewWithSuffix(WordsFile) error = %v", err)
		return
	}
}

func TestNewWithSuffix_Instantiation(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	var words Words
	words, err = NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		words  Words
		suffix core.Suffix
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"check invalid words", args{words: nil, suffix: "_EN"}, core.ErrWordsNil},
		{"check empty suffix", args{words: words, suffix: core.Suffix(internal.Empty)}, core.ErrSuffixIsInvalid},
		{"check invalid suffix", args{words: words, suffix: "  "}, core.ErrSuffixIsInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewWithSuffix(tt.args.words, tt.args.suffix); got == nil {
				t.Errorf("NewWithSuffix() got nil error, want = %v", tt.want)
			}
		})
	}
}

func TestWithSuffix_Get(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wEN, err := NewWithSuffix(wRepository, "_EN")
	if err != nil {
		t.Fatal(err)
	}
	wFA, err := NewWithSuffix(wRepository, "_FA")
	if err != nil {
		t.Fatal(err)
	}
	wFAsl, err := NewWithSuffix(wRepository, " _FA")
	if err != nil {
		t.Fatal(err)
	}
	wFAsr, err := NewWithSuffix(wRepository, "_FA ")
	if err != nil {
		t.Fatal(err)
	}
	wFAsb, err := NewWithSuffix(wRepository, " _FA ")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name  string
		words Words
		arg   string
		want  string
	}{
		{"found EN", wEN, "k1", "v1 EN"},
		{"found FA", wFA, "k1", "v1 FA"},
		{"found FA space left", wFAsl, "kleft", "v2 FA space left"},
		{"found FA space right", wFAsr, "kright", "v3 FA space right"},
		{"found FA space both", wFAsb, "kboth", "v4 FA space both"},
		{"notfound EN", wEN, key_NOTFOUND, internal.Empty},
		{"notfound FA", wFA, key_NOTFOUND, internal.Empty},
		{"empty EN", wEN, internal.Empty, internal.Empty},
		{"empty FA", wFA, internal.Empty, internal.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.words.Get(tt.arg); got != tt.want {
				t.Errorf("WithSuffix.Get() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestWithSuffix_Find(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wEN, err := NewWithSuffix(wRepository, "_EN")
	if err != nil {
		t.Fatal(err)
	}
	wFA, err := NewWithSuffix(wRepository, "_FA")
	if err != nil {
		t.Fatal(err)
	}
	wFAsl, err := NewWithSuffix(wRepository, " _FA")
	if err != nil {
		t.Fatal(err)
	}
	wFAsr, err := NewWithSuffix(wRepository, "_FA ")
	if err != nil {
		t.Fatal(err)
	}
	wFAsb, err := NewWithSuffix(wRepository, " _FA ")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name  string
		words Words
		arg   string
		want  string
		found bool
	}{
		{"found EN", wEN, "k1", "v1 EN", true},
		{"found FA", wFA, "k1", "v1 FA", true},
		{"found FA space left", wFAsl, "kleft", "v2 FA space left", true},
		{"found FA space right", wFAsr, "kright", "v3 FA space right", true},
		{"found FA space both", wFAsb, "kboth", "v4 FA space both", true},
		{"notfound EN", wEN, key_NOTFOUND, internal.Empty, false},
		{"notfound FA", wFA, key_NOTFOUND, internal.Empty, false},
		{"empty EN", wEN, internal.Empty, internal.Empty, false},
		{"empty FA", wFA, internal.Empty, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := tt.words.Find(tt.arg)
			if got != tt.want {
				t.Errorf("WithSuffix.Find() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("WithSuffix.Find() found = %v, want %v", found, tt.found)
			}
		})
	}
}
