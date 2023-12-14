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

func TestWordsCollection_Words_InterfaceSatisfaction(t *testing.T) {
	var _ Words = WordsCollection{}
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
