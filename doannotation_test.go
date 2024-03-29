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

func TestNewDoAnnotation(t *testing.T) {
	// Arrange
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	// Act
	_, err = NewDoAnnotation(wRepository)
	// Assert
	if err != nil {
		t.Errorf("NewDoAnnotation() error = %v", err)
	}
}

func TestNewDoAnnotation_Instantiation(t *testing.T) {
	tests := []struct {
		name  string
		words Words
		want  error
	}{
		{"check invalid words", nil, core.ErrWordsNil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := NewDoAnnotation(tt.words); got == nil {
				t.Errorf("NewDoAnnotation() got nil error, want = %v", tt.want)
			}
		})
	}
}

func TestDoAnnotation_GetNamed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments map[string]any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "k1", arguments: map[string]any{"value": "v1"}}, "v1"},
		{"multiple", args{name: "k2", arguments: map[string]any{"value": "v2", "num": 2}}, "v2,2"},
		{"types", args{name: "k_types", arguments: map[string]any{"string": "x", "integer": 1, "float": 2.2, "boolean": true}}, "x,1,2.2,true"},
		{"person", args{name: "person", arguments: map[string]any{"name": "Saleh", "age": 38, "language": "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
		{"nil annotation arguments", args{name: "k1", arguments: nil}, "{{value}}"},
		{"empty annotation argument", args{name: "k1", arguments: map[string]any{}}, "{{value}}"},
		{"empty annotation", args{name: "k3", arguments: map[string]any{}}, "v3"},
		{"empty value", args{name: "k4", arguments: map[string]any{"value": "v1"}}, internal.Empty},
		{"absent annotation", args{name: "k1", arguments: map[string]any{key_NOTFOUND: "v1"}}, "{{value}}"},
		{"few annotations", args{name: "k2", arguments: map[string]any{"value": "v1"}}, "v1,{{num}}"},
		{"notfound", args{name: key_NOTFOUND, arguments: map[string]any{"value": "v1"}}, internal.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.GetNamed(tt.args.name, tt.args.arguments); got != tt.want {
				t.Errorf("DoAnnotation.GetNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoAnnotation_FindNamed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments map[string]any
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "k1", arguments: map[string]any{"value": "v1"}}, "v1", true},
		{"multiple", args{name: "k2", arguments: map[string]any{"value": "v2", "num": 2}}, "v2,2", true},
		{"types", args{name: "k_types", arguments: map[string]any{"string": "x", "integer": 1, "float": 2.2, "boolean": true}}, "x,1,2.2,true", true},
		{"nil annotation arguments", args{name: "k1", arguments: nil}, "{{value}}", true},
		{"empty annotation argument", args{name: "k1", arguments: map[string]any{}}, "{{value}}", true},
		{"empty annotation", args{name: "k3", arguments: map[string]any{}}, "v3", true},
		{"empty value", args{name: "k4", arguments: map[string]any{"value": "v1"}}, internal.Empty, true},
		{"absent annotation", args{name: "k1", arguments: map[string]any{key_NOTFOUND: "v1"}}, "{{value}}", true},
		{"few annotations", args{name: "k2", arguments: map[string]any{"value": "v1"}}, "v1,{{num}}", true},
		{"notfound", args{name: key_NOTFOUND, arguments: map[string]any{"value": "v1"}}, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.FindNamed(tt.args.name, tt.args.arguments)
			if got != tt.want {
				t.Errorf("DoAnnotation.FindNamed() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("DoAnnotation.FindNamed() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestDoAnnotation_GetIndexed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "kindexed1", arguments: []any{"v1"}}, "v1"},
		{"multiple", args{name: "kindexed2", arguments: []any{"1", 2}}, "1,2"},
		{"types", args{name: "kindexed_types", arguments: []any{"x", 1, 2.2, true}}, "x,1,2.2,true"},
		{"disordered", args{name: "kindexed2_disordered", arguments: []any{"1", 2, "3", 4}}, "2,1,4,3"},
		{"over index", args{name: "kindexed1", arguments: []any{"1", "2", "3"}}, "1"},
		{"few index", args{name: "kindexed2", arguments: []any{"1"}}, "1,{{2}}"},
		{"person index", args{name: "personindexed", arguments: []any{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
		{"nonindex", args{name: "kindexed_nonindex", arguments: []any{"v1"}}, "{{value}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.GetIndexed(tt.args.name, tt.args.arguments...); got != tt.want {
				t.Errorf("DoAnnotation.GetIndexed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoAnnotation_FindIndexed(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []any
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "kindexed1", arguments: []any{"v1"}}, "v1", true},
		{"multiple", args{name: "kindexed2", arguments: []any{"1", 2}}, "1,2", true},
		{"types", args{name: "kindexed_types", arguments: []any{"x", 1, 2.2, true}}, "x,1,2.2,true", true},
		{"disordered", args{name: "kindexed2_disordered", arguments: []any{"1", 2, "3", 4}}, "2,1,4,3", true},
		{"over index", args{name: "kindexed1", arguments: []any{"1", "2", "3"}}, "1", true},
		{"few index", args{name: "kindexed2", arguments: []any{"1"}}, "1,{{2}}", true},
		{"person index", args{name: "personindexed", arguments: []any{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer.", true},
		{"nonindex", args{name: "kindexed_nonindex", arguments: []any{"v1"}}, "{{value}}", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.FindIndexed(tt.args.name, tt.args.arguments...)
			if got != tt.want {
				t.Errorf("DoAnnotation.FindIndexed() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("DoAnnotation.FindIndexed() found = %v, want %v", found, tt.found)
			}
		})
	}
}

func TestDoAnnotation_GetFormatted(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{name: "kformatted1", arguments: []any{"v1"}}, "v1"},
		{"multiple", args{name: "kformatted2", arguments: []any{"1", 2}}, "1,2"},
		{"types", args{name: "kformatted_types", arguments: []any{"x", 1, 2.2, true}}, "x,1,2.200000,true"},
		{"person", args{name: "personformatted", arguments: []any{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wAnnotation.GetFormatted(tt.args.name, tt.args.arguments...); got != tt.want {
				t.Errorf("DoAnnotation.GetFormatted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoAnnotation_FindFormatted(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "doannotation"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	wAnnotation, err := NewDoAnnotation(wRepository)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		name      string
		arguments []any
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"single", args{name: "kformatted1", arguments: []any{"v1"}}, "v1", true},
		{"multiple", args{name: "kformatted2", arguments: []any{"1", 2}}, "1,2", true},
		{"types", args{name: "kformatted_types", arguments: []any{"x", 1, 2.2, true}}, "x,1,2.200000,true", true},
		{"person", args{name: "personformatted", arguments: []any{"Saleh", 38, "Golang"}}, "Hi, my name is Saleh, I'm 38 years old and a Golang developer.", true},
		{"notfound", args{name: key_NOTFOUND, arguments: []any{}}, internal.Empty, false},
		{"empty", args{name: "k4", arguments: []any{}}, internal.Empty, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := wAnnotation.FindFormatted(tt.args.name, tt.args.arguments...)
			if got != tt.want {
				t.Errorf("DoAnnotation.FindFormatted() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("DoAnnotation.FindFormatted() found = %v, want %v", found, tt.found)
			}
		})
	}
}

//┌ Benchmark
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func BenchmarkDoAnnotationNamed(b *testing.B) {
	source := `k1=First {{first}} , Second {{second}} , Third {{third}}`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewDoAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.FindNamed("k1", map[string]any{
			"first":  1,
			"second": 2.22,
			"three":  "three",
		})
	}
}

func BenchmarkDoAnnotationIndexed(b *testing.B) {
	source := `k1=First {{1}} , Second {{2}} , Third {{3}}`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewDoAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.FindIndexed("k1", 1, 2.22, "three")
	}
}

func BenchmarkDoAnnotationFormatted(b *testing.B) {
	source := `k1=First %s , Second %s , Third %s`
	wCollection, err := NewWordsCollection(string(source), core.Separator, core.Comment)
	if err != nil {
		b.Fatal(err)
	}
	w, err := NewDoAnnotation(wCollection)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.FindFormatted("k1", 1, 2.22, "three")
	}
}
