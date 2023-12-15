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

func TestWordsCollection_Words_InterfaceSatisfaction(t *testing.T) {
	var _ Words = WordsCollection{}
}

func TestNewWordsCollection(t *testing.T) {
	// Arrange
	source, err := os.ReadFile(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	// Act
	_, err = NewWordsCollection(string(source), core.Separator, core.Comment)
	// Assert
	if err != nil {
		t.Errorf("NewWordsCollection() error = %v", err)
		return
	}
}

func TestNewWordsCollection_Instantiation(t *testing.T) {
	valid_source, err := os.ReadFile(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	invalid_absent_name, _ := os.ReadFile(path.Join(path_WORDS, "invalid_absent_name"))
	data_duplicated, _ := os.ReadFile(path.Join(path_WORDS, "collection_duplicate"))
	type args struct {
		source    string
		separator rune
		comment   rune
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"check invalid source", args{ source: internal.Empty, separator: core.Separator, comment: core.Comment, }, core.ErrWordsEmpty},
		{"check invalid separator delimiters", args{ source: string(valid_source), separator: 'x', comment: core.Comment, }, core.ErrSeparatorIsInvalid},
		{"check invalid comment delimiters", args{ source: string(valid_source), separator: core.Separator, comment: 'x', }, core.ErrCommentIsInvalid},
		{"check invalid normalization", args{ source: string(invalid_absent_name), separator: core.Separator, comment: core.Comment, }, core.ErrNameNotPresent},
		{"check invalid treasure", args{ source: string(data_duplicated), separator: core.Separator, comment: core.Comment, }, core.ErrNameDuplicated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewWordsCollection(tt.args.source, tt.args.separator, tt.args.comment); got == nil {
				t.Errorf("NewWordsCollection() got nil error, want %v", tt.want)
			}
		})
	}
}

func TestWordsCollection_Get(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	w, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsCollection() error = %v", err)
	}
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"found", "k1", "v1"},
		{"notfound", "kx", internal.Empty},
		{"empty", internal.Empty, internal.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := w.Get(tt.arg); got != tt.want {
				t.Errorf("WordsCollection.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordsCollection_Find(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "valid_sparse__want"))
	if err != nil {
		t.Fatal(err)
	}
	w, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsCollection() error = %v", err)
	}
	tests := []struct {
		name  string
		arg   string
		want  string
		found bool
	}{
		{"found", "k1", "v1", true},
		{"found empty", "k11", internal.Empty, true},
		{"notfound", "kx", internal.Empty, false},
		{"empty", internal.Empty, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := w.Find(tt.arg)
			if got != tt.want {
				t.Errorf("WordsCollection.Find() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("WordsCollection.Find() found = %v, want %v", found, tt.found)
			}
		})
	}
}

//┌ Benchmark
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func BenchmarkWordsCollection(b *testing.B) {
	source, err := os.ReadFile(path.Join(path_BENCHMARK, "normalization__large"))
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, found := w.Find("k1000")
		if !found {
			b.Fatal("error not found")
		}
	}
}
