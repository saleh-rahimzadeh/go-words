package gowords_test

import (
	"os"
	"path"
	"testing"

	. "github.com/saleh-rahimzadeh/go-words"

	"github.com/saleh-rahimzadeh/go-words/core"
	"github.com/saleh-rahimzadeh/go-words/internal"
)

//┌ Tests
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func TestGetBy(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)
	type args struct {
		w      Words
		name   string
		suffix core.Suffix
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"found EN", args{wRepository, "k1", EN}, "v1 EN"},
		{"found FA", args{wRepository, "k1", FA}, "v1 FA"},
		{"found FA space left", args{wRepository, "kleft", " " + FA}, "v2 FA space left"},
		{"found FA space right", args{wRepository, "kright", FA + " "}, "v3 FA space right"},
		{"found FA space both", args{wRepository, "kboth", " " + FA + " "}, "v4 FA space both"},
		{"notfound EN", args{wRepository, key_NOTFOUND, EN}, internal.Empty},
		{"notfound FA", args{wRepository, key_NOTFOUND, FA}, internal.Empty},
		{"empty EN", args{wRepository, internal.Empty, EN}, internal.Empty},
		{"empty FA", args{wRepository, internal.Empty, FA}, internal.Empty},
		{"empty suffix", args{wRepository, "k1", core.Suffix(internal.Empty)}, internal.Empty},
		{"invalid suffix", args{wRepository, "k1", "  "}, internal.Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBy(tt.args.w, tt.args.name, tt.args.suffix); got != tt.want {
				t.Errorf("GetBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBy(t *testing.T) {
	source, err := os.ReadFile(path.Join(path_WORDS, "withsuffix"))
	if err != nil {
		t.Fatal(err)
	}
	wRepository, err := NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		t.Fatal(err)
	}
	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)
	type args struct {
		w      Words
		name   string
		suffix core.Suffix
	}
	tests := []struct {
		name  string
		args  args
		want  string
		found bool
	}{
		{"found EN", args{wRepository, "k1", EN}, "v1 EN", true},
		{"found FA", args{wRepository, "k1", FA}, "v1 FA", true},
		{"found FA space left", args{wRepository, "kleft", " " + FA}, "v2 FA space left", true},
		{"found FA space right", args{wRepository, "kright", FA + " "}, "v3 FA space right", true},
		{"found FA space both", args{wRepository, "kboth", " " + FA + " "}, "v4 FA space both", true},
		{"notfound EN", args{wRepository, key_NOTFOUND, EN}, internal.Empty, false},
		{"notfound FA", args{wRepository, key_NOTFOUND, FA}, internal.Empty, false},
		{"empty EN", args{wRepository, internal.Empty, EN}, internal.Empty, false},
		{"empty FA", args{wRepository, internal.Empty, FA}, internal.Empty, false},
		{"empty suffix", args{wRepository, "k1", core.Suffix(internal.Empty)}, internal.Empty, false},
		{"invalid suffix", args{wRepository, "k1", "  "}, internal.Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := FindBy(tt.args.w, tt.args.name, tt.args.suffix)
			if got != tt.want {
				t.Errorf("FindBy() got = %v, want %v", got, tt.want)
			}
			if found != tt.found {
				t.Errorf("FindBy() found = %v, want %v", found, tt.found)
			}
		})
	}
}
