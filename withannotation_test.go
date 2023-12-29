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

func TestNewWithAnnotation(t *testing.T) {
	// Arrange
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	// Act
	_, err = NewWithAnnotation(wRepository)
	// Assert
	if err != nil {
		t.Errorf("NewWithAnnotation() error = %v", err)
	}
}

func TestNewWithAnnotation_Instantiation(t *testing.T) {
	tests := []struct {
		name  string
		words Words
		want  error
	}{
		{"check invalid words", nil, core.ErrWordsNil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewWithAnnotation(tt.words); got == nil {
				t.Errorf("NewWithAnnotation() got nil error, want = %v", tt.want)
			}
		})
	}
}

func TestWithAnnotation_Get(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "k1", arguments: map[string]string{"value": "v1"}}, "v1"},
		{"multiple", args{name: "k2", arguments: map[string]string{"value": "v2", "num": "2"}}, "v2,2"},
		{"person", args{name: "person", arguments: map[string]string{"name": "Saleh", "age": "38", "language": "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
		{"nil annotation arguments", args{name: "k1", arguments: nil}, "{{value}}"},
		{"empty annotation argument", args{name: "k1", arguments: map[string]string{}}, "{{value}}"},
		{"empty annotation", args{name: "k3", arguments: map[string]string{}}, "v3"},
		{"empty value", args{name: "k4", arguments: map[string]string{"value": "v1"}}, internal.Empty},
		{"absent annotation", args{name: "k1", arguments: map[string]string{key_NOTFOUND: "v1"}}, "{{value}}"},
		{"notfound", args{name: key_NOTFOUND, arguments: map[string]string{"value": "v1"}}, internal.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.Get(tt.args.name, tt.args.arguments); got != tt.want {
				t.Errorf("WithAnnotation.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAnnotation_Find(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments map[string]string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "k1", arguments: map[string]string{"value": "v1"}}, "v1", true},
		{"multiple", args{name: "k2", arguments: map[string]string{"value": "v2", "num": "2"}}, "v2,2", true},
		{"nil annotation arguments", args{name: "k1", arguments: nil}, "{{value}}", true},
		{"empty annotation argument", args{name: "k1", arguments: map[string]string{}}, "{{value}}", true},
		{"empty annotation", args{name: "k3", arguments: map[string]string{}}, "v3", true},
		{"empty value", args{name: "k4", arguments: map[string]string{"value": "v1"}}, internal.Empty, true},
		{"absent annotation", args{name: "k1", arguments: map[string]string{key_NOTFOUND: "v1"}}, "{{value}}", true},
		{"notfound", args{name: key_NOTFOUND, arguments: map[string]string{"value": "v1"}}, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.Find(tt.args.name, tt.args.arguments)
			if got != tt.want {
				t.Errorf("WithAnnotation.Find() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("WithAnnotation.Find() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestWithAnnotation_GetIndexed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "kindexed1", arguments: []string{"v1"}}, "v1"},
		{"multiple", args{name: "kindexed2", arguments: []string{"1", "2"}}, "1,2"},
		{"person index", args{name: "personindexed", arguments: []string{"Saleh", "38", "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.GetIndexed(tt.args.name, tt.args.arguments...); got != tt.want {
				t.Errorf("WithAnnotation.GetIndexed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAnnotation_FindIndexed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "kindexed1", arguments: []string{"v1"}}, "v1", true},
		{"multiple", args{name: "kindexed2", arguments: []string{"1", "2"}}, "1,2", true},
		{"person index", args{name: "personindexed", arguments: []string{"Saleh", "38", "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer.", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.FindIndexed(tt.args.name, tt.args.arguments...)
			if got != tt.want {
				t.Errorf("WithAnnotation.FindIndexed() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("WithAnnotation.FindIndexed() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestWithAnnotation_GetFormatted(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "kformatted1", arguments: []interface{}{"v1"}}, "v1"},
		{"multiple", args{name: "kformatted2", arguments: []interface{}{"1", 2}}, "1,2"},
		{"person index", args{name: "personformatted", arguments: []interface{}{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.GetFormatted(tt.args.name, tt.args.arguments...); got != tt.want {
				t.Errorf("WithAnnotation.GetFormatted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithAnnotation_FindFormatted(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewWithAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "kformatted1", arguments: []interface{}{"v1"}}, "v1", true},
		{"multiple", args{name: "kformatted2", arguments: []interface{}{"1", 2}}, "1,2", true},
		{"person index", args{name: "personformatted", arguments: []interface{}{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer.", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.FindFormatted(tt.args.name, tt.args.arguments...)
			if got != tt.want {
				t.Errorf("WithAnnotation.FindFormatted() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("WithAnnotation.FindFormatted() found = %v, want %v", found, tt.found)
			}
		})
	}
}

//┌ Benchmark
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func BenchmarkWithAnnotation(b *testing.B) {
	source := `k1=First {{first}} , Second {{second}} , Third {{third}}`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewWithAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Find("k1", map[string]string{
			"first":  "one",
			"second": "two",
			"three":  "three",
		})
	}
}

func BenchmarkWithAnnotationIndexed(b *testing.B) {
	source := `k1=First {{1}} , Second {{2}} , Third {{3}}`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewWithAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.FindIndexed("k1", "one", "two", "three")
	}
}

func BenchmarkWithAnnotationFormatted(b *testing.B) {
	source := `k1=First %s , Second %s , Third %s`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewWithAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.FindFormatted("k1", "one", "two", "three")
	}
}
