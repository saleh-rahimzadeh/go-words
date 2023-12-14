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

func TestNewWordsFile(t *testing.T) {
	// Arrange
	file, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	// Act
	_, err = NewWordsFile(file, core.Separator, core.Comment)
	// Assert
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
}

func TestWordsFile_Words_InterfaceSatisfaction(t *testing.T) {
	var _ Words = WordsFile{}
}

func TestWordsFile_CheckError(t *testing.T) {
	fileValid, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileValid.Close()
	wValid, err := NewWordsFile(fileValid, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	fileDuplicate, err := os.Open(path.Join(path_WORDS, "duplicate_found"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileDuplicate.Close()
	wDuplicate, err := NewWordsFile(fileDuplicate, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	fileAbsentName, err := os.Open(path.Join(path_WORDS, "invalid_absent_name"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileAbsentName.Close()
	wAbsentName, err := NewWordsFile(fileAbsentName, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	fileNoSeparator, err := os.Open(path.Join(path_WORDS, "invalid_no_separator"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileNoSeparator.Close()
	wNoSeparator, err := NewWordsFile(fileNoSeparator, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		words   *WordsFile
		wantErr bool
	}{
		{"valid", &wValid, false},
		{"invalid_duplicate", &wDuplicate, true},
		{"invalid_absent_name", &wAbsentName, true},
		{"invalid_no_separator", &wNoSeparator, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.words.CheckError(); (err != nil) != tt.wantErr {
				t.Errorf("WordsFile.CheckError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWordsFile_Get(t *testing.T) {
	file, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
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
				t.Errorf("WordsFile.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordsFile_Find(t *testing.T) {
	file, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	tests := []struct {
		name      string
		arg       string
		wantValue string
		wantFound bool
	}{
		{"found", "k1", "v1", true},
		{"notfound", "kx", internal.Empty, false},
		{"empty", internal.Empty, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := w.Find(tt.arg)
			if gotValue != tt.wantValue {
				t.Errorf("WordsFile.Find() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("WordsFile.Find() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

//┌ Benchmark
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func BenchmarkWordsFile(b *testing.B) {
	file, err := os.Open(path.Join(path_BENCHMARK, "normalization__large"))
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, found := w.Find("k1000")
		if err := w.Err(); err != nil {
			b.Fatal(err)
		}
		if !found {
			b.Fatal("error not found")
		}
	}
}

func BenchmarkWordsFileUnsafe(b *testing.B) {
	file, err := os.Open(path.Join(path_BENCHMARK, "normalization__large"))
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, found := w.FindUnsafe("k1000")
		if err := w.Err(); err != nil {
			b.Fatal(err)
		}
		if !found {
			b.Fatal("error not found")
		}
	}
}
