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

func TestNewWordsFile_Instantiation(t *testing.T) {
	file, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	type args struct {
		file      *os.File
		separator rune
		comment   rune
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{"check nil file", args{file: nil, separator: core.Separator, comment: core.Comment}, core.ErrNilFile},
		{"check empty file", args{file: &os.File{}, separator: core.Separator, comment: core.Comment}, core.ErrNilFile},
		{"check invalid separator delimiters", args{file: file, separator: 'x', comment: core.Comment}, core.ErrSeparatorIsInvalid},
		{"check invalid comment delimiters", args{file: file, separator: core.Separator, comment: 'x'}, core.ErrCommentIsInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewWordsFile(tt.args.file, tt.args.separator, tt.args.comment); got == nil {
				t.Errorf("NewWordsFile() got nil error, want = %v", tt.want)
			}
		})
	}
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
	fileValidSparse, err := os.Open(path.Join(path_WORDS, "valid_sparse__source"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileValidSparse.Close()
	wSparse, err := NewWordsFile(fileValidSparse, core.Separator, core.Comment)
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
		{"valid_sparse", &wSparse, false},
		{"invalid_duplicate", &wDuplicate, true},
		{"invalid_absent_name", &wAbsentName, true},
		{"invalid_no_separator", &wNoSeparator, true},
		{"empty WordsFile", &WordsFile{}, true},
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
		{"notfound", key_NOTFOUND, internal.Empty},
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
		{"notfound", key_NOTFOUND, internal.Empty, false},
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

func TestWordsFile_FindUnsafe(t *testing.T) {
	file, err := os.Open(path.Join(path_WORDS, "valid_sparse__source"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	const (
		wantValue string = "v7"
		wantFound bool   = true
	)
	gotValue, gotFound := w.FindUnsafe("k7")
	gotError := w.Err()
	if gotError != nil {
		t.Errorf("WordsFile.FindUnsafe() error = %v", gotError)
		return
	}
	if gotValue != wantValue {
		t.Errorf("WordsFile.FindUnsafe() gotValue = %v, want %v", gotValue, wantValue)
		return
	}
	if gotFound != wantFound {
		t.Errorf("WordsFile.FindUnsafe() gotFound = %v, want %v", gotFound, wantFound)
		return
	}
}

func TestWordsFile_FindUnsafe_Invalid(t *testing.T) {
	file, err := os.Open(path.Join(path_WORDS, "invalid_absent_name"))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	w, err := NewWordsFile(file, core.Separator, core.Comment)
	if err != nil {
		t.Errorf("NewWordsFile() error = %v", err)
		return
	}
	w.FindUnsafe(key_NOTFOUND)
	if w.Err() == nil {
		t.Errorf("WordsFile.FindUnsafe() got nil error, want = %v", core.ErrNameNotPresent)
		return
	}
}

func TestWordsFile_FindUnsafe_Panic(t *testing.T) {
	_, wfile, _ := os.Pipe()
	wNonSeek, _ := NewWordsFile(wfile, core.Separator, core.Comment)
	closedfile, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	wClosedFile, _ := NewWordsFile(closedfile, core.Separator, core.Comment)
	closedfile.Close()
	tests := []struct {
		name  string
		words *WordsFile
	}{
		{"empty WordsFile", &WordsFile{}},
		{"unsupport seek", &wNonSeek},
		{"closed file", &wClosedFile},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.words.FindUnsafe(key_NOTFOUND)
			if tt.words.Err() == nil {
				t.Errorf("WordsFile.FindUnsafe() got nil error")
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
			b.Fatal(benchmark_KEY_NOTFOUND)
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
			b.Fatal(benchmark_KEY_NOTFOUND)
		}
	}
}
