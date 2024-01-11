package gowords_test

import (
	"os"

	. "github.com/saleh-rahimzadeh/go-words"
)

const (
	path_WORDS     string = "testdata/words/"
	path_BENCHMARK string = "testdata/benchmark/"
)

const (
	key_NOTFOUND           string = "NOT_FOUND_KEY"
	benchmark_KEY_NOTFOUND string = "error: key not found, please generate words for benchmarking by `make generate` command"
)

func init() {
	var _ Words = WordsCollection{}
	var _ Words = WordsFile{}
	var _ Words = WordsRepository{}
	var _ Words = WithSuffix{}
}

func init() {
	var err error
	_, err = os.Stat(path_WORDS)
	if os.IsNotExist(err) {
		panic(err)
	}
}
