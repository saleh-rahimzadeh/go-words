package internal_test

import (
	"os"
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/saleh-rahimzadeh/go-words/core"
	. "github.com/saleh-rahimzadeh/go-words/internal"
)

const (
	path_WORDS     string = "../testdata/words/"
	path_BENCHMARK string = "../testdata/benchmark/"
)

func init() {
	var err error
	_, err = os.Stat(path_WORDS)
	if os.IsNotExist(err) {
		panic(err)
	}
	_, err = os.Stat(path_BENCHMARK)
	if os.IsNotExist(err) {
		panic(err)
	}
}

//┌ Test
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func TestExtract(t *testing.T) {
	type args struct {
		line      string
		name      string
		separator string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantFound bool
	}{
		{"found", args{"k=v", "k", string(core.Separator)}, "v", true},
		{"found empty", args{"k=", "k", string(core.Separator)}, Empty, true},
		{"not found", args{"k=v", "x", string(core.Separator)}, Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := Extract(tt.args.line, tt.args.name, tt.args.separator)
			if gotValue != tt.wantValue {
				t.Errorf("Extract() gotValue = %v, wantValue %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Extract() gotFound = %v, wantFound %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		line      string
		separator string
		comment   string
	}
	tests := []struct {
		name      string
		args      args
		wantKey   string
		wantValue string
		wantErr   bool
	}{
		{"valid", args{"k1=v1", string(core.Separator), string(core.Comment)}, "k1", "v1", false},
		{"valid sparse 2", args{"k2= v2", string(core.Separator), string(core.Comment)}, "k2", "v2", false},
		{"valid sparse 3", args{"k3 =v3", string(core.Separator), string(core.Comment)}, "k3", "v3", false},
		{"valid sparse 4", args{"k4 = v4", string(core.Separator), string(core.Comment)}, "k4", "v4", false},
		{"valid sparse 5", args{"k5      =       v5       ", string(core.Separator), string(core.Comment)}, "k5", "v5", false},
		{"valid sparse 6", args{"     k6      =          v6       ", string(core.Separator), string(core.Comment)}, "k6", "v6", false},
		{"valid sparse 8", args{"k   8   =   v      8", string(core.Separator), string(core.Comment)}, "k   8", "v      8", false},
		{"valid sparse 9", args{"    k     9       =        v        9        ", string(core.Separator), string(core.Comment)}, "k     9", "v        9", false},
		{"empty", args{Empty, string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"empty space", args{" ", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"comment", args{"#comment", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"comment space", args{"# Comment", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"comment empty", args{"#", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"comment double", args{"##", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"different comment", args{"|comment", string(core.Separator), "|"}, Empty, Empty, true},
		{"invalid absent name", args{"=v1", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"invalid incorrect separator", args{"k2$v2", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"invalid no separator", args{"hello", string(core.Separator), string(core.Comment)}, Empty, Empty, true},
		{"different separator", args{"k1|v1", "|", string(core.Comment)}, "k1", "v1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, value, err := Parse(tt.args.line, tt.args.separator, tt.args.comment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if key != tt.wantKey {
				t.Errorf("Parse() got = %v, wantKey %v", key, tt.wantKey)
			}
			if value != tt.wantValue {
				t.Errorf("Parse() got1 = %v, wantValue %v", value, tt.wantValue)
			}
		})
	}
}

func TestNormalizeLine(t *testing.T) {
	type args struct {
		line      string
		separator string
		comment   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"k1=v1", string(core.Separator), string(core.Comment)}, "k1=v1", false},
		{"valid sparse 2", args{"k2= v2", string(core.Separator), string(core.Comment)}, "k2=v2", false},
		{"valid sparse 3", args{"k3 =v3", string(core.Separator), string(core.Comment)}, "k3=v3", false},
		{"valid sparse 4", args{"k4 = v4", string(core.Separator), string(core.Comment)}, "k4=v4", false},
		{"valid sparse 5", args{"k5      =       v5       ", string(core.Separator), string(core.Comment)}, "k5=v5", false},
		{"valid sparse 6", args{"     k6      =          v6       ", string(core.Separator), string(core.Comment)}, "k6=v6", false},
		{"valid sparse 8", args{"k   8   =   v      8", string(core.Separator), string(core.Comment)}, "k   8=v      8", false},
		{"valid sparse 9", args{"    k     9       =        v        9        ", string(core.Separator), string(core.Comment)}, "k     9=v        9", false},
		{"empty", args{Empty, string(core.Separator), string(core.Comment)}, Empty, true},
		{"empty space", args{" ", string(core.Separator), string(core.Comment)}, Empty, true},
		{"comment", args{"#comment", string(core.Separator), string(core.Comment)}, Empty, true},
		{"comment space", args{"# Comment", string(core.Separator), string(core.Comment)}, Empty, true},
		{"comment empty", args{"#", string(core.Separator), string(core.Comment)}, Empty, true},
		{"comment double", args{"##", string(core.Separator), string(core.Comment)}, Empty, true},
		{"different comment", args{"|comment", string(core.Separator), "|"}, Empty, true},
		{"invalid absent name", args{"=v1", string(core.Separator), string(core.Comment)}, Empty, true},
		{"invalid incorrect separator", args{"k2$v2", string(core.Separator), string(core.Comment)}, Empty, true},
		{"invalid no separator", args{"hello", string(core.Separator), string(core.Comment)}, Empty, true},
		{"different separator", args{"k1|v1", "|", string(core.Comment)}, "k1|v1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeLine(tt.args.line, tt.args.separator, tt.args.comment)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NormalizeLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalization(t *testing.T) {
	valid_oneline__source, _ := os.ReadFile(path.Join(path_WORDS, "valid_oneline__source"))
	valid_oneline__want, _ := os.ReadFile(path.Join(path_WORDS, "valid_oneline__want"))
	valid__source, _ := os.ReadFile(path.Join(path_WORDS, "valid__source"))
	valid__want, _ := os.ReadFile(path.Join(path_WORDS, "valid__want"))
	valid_sparse__source, _ := os.ReadFile(path.Join(path_WORDS, "valid_sparse__source"))
	valid_sparse__want, _ := os.ReadFile(path.Join(path_WORDS, "valid_sparse__want"))
	invalid_absent_name, _ := os.ReadFile(path.Join(path_WORDS, "invalid_absent_name"))
	invalid_no_separator, _ := os.ReadFile(path.Join(path_WORDS, "invalid_no_separator"))
	invalid_incorrect_separator, _ := os.ReadFile(path.Join(path_WORDS, "invalid_incorrect_separator"))
	invalid_incorrect_comment, _ := os.ReadFile(path.Join(path_WORDS, "invalid_incorrect_comment"))
	valid_different_separator__source, _ := os.ReadFile(path.Join(path_WORDS, "valid_different_separator__source"))
	valid_different_separator__want, _ := os.ReadFile(path.Join(path_WORDS, "valid_different_separator__want"))
	valid_different_comment__source, _ := os.ReadFile(path.Join(path_WORDS, "valid_different_comment__source"))
	valid_different_comment__want, _ := os.ReadFile(path.Join(path_WORDS, "valid_different_comment__want"))
	type args struct {
		source    string
		separator string
		comment   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid oneline", args{string(valid_oneline__source), string(core.Separator), string(core.Comment)}, string(valid_oneline__want), false},
		{"valid", args{string(valid__source), string(core.Separator), string(core.Comment)}, string(valid__want), false},
		{"valid sparse", args{string(valid_sparse__source), string(core.Separator), string(core.Comment)}, string(valid_sparse__want), false},
		{"invalid absent name", args{string(invalid_absent_name), string(core.Separator), string(core.Comment)}, Empty, true},
		{"invalid no separator", args{string(invalid_no_separator), string(core.Separator), string(core.Comment)}, Empty, true},
		{"invalid incorrect separator", args{string(invalid_incorrect_separator), string(core.Separator), string(core.Comment)}, Empty, true},
		{"invalid incorrect comment", args{string(invalid_incorrect_comment), string(core.Separator), string(core.Comment)}, Empty, true},
		{"valid different separator", args{string(valid_different_separator__source), string('|'), string(core.Comment)}, string(valid_different_separator__want), false},
		{"valid different comment", args{string(valid_different_comment__source), string(core.Separator), string('|')}, string(valid_different_comment__want), false},
		{"empty", args{Empty, string(core.Separator), string(core.Comment)}, Empty, false},
		{"whitespace", args{"      ", string(core.Separator), string(core.Comment)}, Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Normalization(tt.args.source, tt.args.separator, tt.args.comment)
			if (err != nil) != tt.wantErr {
				t.Errorf("Normalization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			desire := strings.Join(got, NewLine)
			if desire != tt.want {
				t.Errorf("Normalization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidationSource(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{"source empty", Empty, true},
		{"source empty space", "   ", true},
		{"source empty line", "  \n  \n    ", true},
		{"valid", "-", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidationSource(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("ValidationSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidationDelimiters(t *testing.T) {
	type args struct {
		separator string
		comment   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"same separator and comment", args{string(core.Separator), string(core.Separator)}, true},
		{"invalid separator", args{string('x'), string(core.Comment)}, true},
		{"invalid separator space", args{string(' '), string(core.Comment)}, true},
		{"invalid comment", args{string(core.Separator), string('x')}, true},
		{"invalid comment space", args{string(core.Separator), string(' ')}, true},
		{"valid", args{string(core.Separator), string(core.Comment)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidationDelimiters(tt.args.separator, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("ValidationDelimiters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckDuplication(t *testing.T) {
	duplicate_nofound, _ := os.ReadFile(path.Join(path_WORDS, "duplicate_nofound"))
	duplicate_found, _ := os.ReadFile(path.Join(path_WORDS, "duplicate_found"))
	type args struct {
		source    []string
		separator string
	}
	tests := []struct {
		name     string
		args     args
		want     bool
		wantName string
	}{
		{"duplicate notfound", args{strings.Split(string(duplicate_nofound), NewLine), string(core.Separator)}, false, Empty},
		{"duplicate found", args{strings.Split(string(duplicate_found), NewLine), string(core.Separator)}, true, "k2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFound, gotKey := CheckDuplication(tt.args.source, tt.args.separator)
			if gotFound != tt.want {
				t.Errorf("CheckDuplication() gotFound = %v, want %v", gotFound, tt.want)
			}
			if gotKey != tt.wantName {
				t.Errorf("CheckDuplication() gotKey = %v, want %v", gotKey, tt.wantName)
			}
		})
	}
}

func TestValidationName(t *testing.T) {
	tests := []struct {
		name        string
		argName     string
		wantName    string
		wantIsValid bool
	}{
		{"valid", "k1", "k1", true},
		{"valid space", "k 1", "k 1", true},
		{"valid trim", "  k1  ", "k1", true},
		{"invalid", "k\n1", Empty, false},
		{"empty", Empty, Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotIsValid := ValidationName(tt.argName)
			if gotName != tt.wantName {
				t.Errorf("ValidationName() gotName = %v, wantName %v", gotName, tt.wantName)
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("ValidationName() gotIsValid = %v, wantIsValid %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestTreasure(t *testing.T) {
	data_valid, _ := os.ReadFile(path.Join(path_WORDS, "collection"))
	data_duplicated, _ := os.ReadFile(path.Join(path_WORDS, "collection_duplicate"))
	tests := []struct {
		name    string
		source  []string
		want    map[string]string
		wantErr bool
	}{
		{"valid", strings.Split(string(data_valid), NewLine), map[string]string{
			"k1":      "v1",
			"k2":      "v2",
			"k 3":     "v 3",
			"k     4": "v        4",
			"k5":      "v5",
			"k6":      "",
			"k7":      "v7",
		}, false},
		{"duplicated", strings.Split(string(data_duplicated), NewLine), nil, true},
	}
	var separator = string(core.Separator)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Treasure(tt.source, separator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Treasure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Treasure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidationSuffix(t *testing.T) {
	tests := []struct {
		name        string
		arg         string
		wantSuffix  string
		wantIsValid bool
	}{
		{"valid", "EN", "EN", true},
		{"valid space", " EN", " EN", true},
		{"valid trim", " EN ", " EN", true},
		{"invalid", "EN\nEN", Empty, false},
		{"space", "  ", Empty, false},
		{"empty", Empty, Empty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotIsValid := ValidationSuffix(tt.arg)
			if gotName != tt.wantSuffix {
				t.Errorf("ValidationSuffix() gotName = %v, wantName %v", gotName, tt.wantSuffix)
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("ValidationSuffix() gotIsValid = %v, wantIsValid %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestValidationFile(t *testing.T) {
	fileValid, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	defer fileValid.Close()
	if err := ValidationFile(fileValid); err != nil {
		t.Errorf("TestValidationFile() error = %v", err)
		return
	}
	fileEmpty, err := os.CreateTemp("", "gowords_TestValidationFile_empty_file")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fileEmpty.Name())
	defer fileEmpty.Close()
	fileClosed, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		t.Fatal(err)
	}
	fileClosed.Close()
	tests := []struct {
		name string
		file *os.File
		want error
	}{
		{"check nil file", nil, core.ErrFileNil},
		{"check zero instance file", &os.File{}, core.ErrFileNil},
		{"check closed file", fileClosed, (os.PathError{}).Err},
		{"check empty file", fileEmpty, core.ErrFileEmpty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidationFile(tt.file); got == nil {
				t.Errorf("ValidationFile() got nil error, want %v", tt.want)
			}
		})
	}
}

//┌ Benchmark
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func BenchmarkNormalizationSmall(b *testing.B) {
	data, err := os.ReadFile(path.Join(path_BENCHMARK, "normalization__small"))
	if err != nil {
		b.Fatal(err)
	}
	var source string = string(data)
	var separator string = string(core.Separator)
	var comment string = string(core.Comment)
	for i := 0; i < b.N; i++ {
		Normalization(source, separator, comment)
	}
}

func BenchmarkNormalizationLarge(b *testing.B) {
	data, err := os.ReadFile(path.Join(path_BENCHMARK, "normalization__large"))
	if err != nil {
		b.Fatal(err)
	}
	var source string = string(data)
	var separator string = string(core.Separator)
	var comment string = string(core.Comment)
	for i := 0; i < b.N; i++ {
		Normalization(source, separator, comment)
	}
}
