package gowords_test

import (
	"os"

	. "github.com/saleh-rahimzadeh/go-words"
)

const (
	path_WORDS     string = "testdata/words/"
	path_BENCHMARK string = "testdata/benchmark/"
)

func init() {
	var _ Words = WordsCollection{}
	var _ Words = WordsFile{}
	var _ Words = WordsRepository{}
}

func init() {
	var err error
	_, err = os.Stat(path_WORDS)
	if os.IsNotExist(err) {
		panic(err)
	}
}
