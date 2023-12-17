package gowords_test

import (
	"fmt"
	"os"
	"path"

	gowords "github.com/saleh-rahimzadeh/go-words"
)

//┌ WordsRepository Examples
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleWordsRepository() {
	const separator rune = '='
	const comment rune = '#'
	const source = `
k1=v1
k2=v2
# comment
k3=v3
`

	w, err := gowords.NewWordsRepository(source, separator, comment)
	if err != nil {
		panic(err)
	}

	value := w.Get("k1")
	fmt.Println(value)

	//Output: v1
}

func ExampleWordsRepository_Find() {
	const separator rune = '='
	const comment rune = '#'
	const source = `
k1=v1
k2=v2
k3=v3
`

	w, err := gowords.NewWordsRepository(source, separator, comment)
	if err != nil {
		panic(err)
	}

	value, found := w.Find("k1")
	if found {
		fmt.Println(value)
	}

	//Output: v1
}

func ExampleWordsRepository_customSeparator() {
	const separator rune = ':'
	const comment rune = '#'
	const source string = `
k1:v1
k2:v2
k3:v3
`

	w, err := gowords.NewWordsRepository(source, separator, comment)
	if err != nil {
		panic(err)
	}

	value := w.Get("k1")
	fmt.Println(value)

	//Output: v1
}

func ExampleWordsRepository_customComment() {
	const separator rune = '='
	const comment rune = '@'
	const source string = `
k1=v1
@ this is a comment
k2=v2
`

	w, err := gowords.NewWordsRepository(source, separator, comment)
	if err != nil {
		panic(err)
	}

	value := w.Get("k1")
	fmt.Println(value)

	//Output: v1
}

//┌ WordsCollection Examples
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleWordsCollection() {
	const separator rune = '='
	const comment rune = '#'
	const source = `
k1=v1
k2=v2
k3=v3
`

	w, err := gowords.NewWordsCollection(source, separator, comment)
	if err != nil {
		panic(err)
	}

	value_1, found := w.Find("k1")
	if found {
		fmt.Println(value_1)
	}

	value_2 := w.Get("k1")
	fmt.Println(value_2)

	//Output:
	// v1
	// v1
}

//┌ WordsFile Examples
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleWordsFile() {
	file, err := os.Open(path.Join(path_WORDS, "valid__want"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const separator rune = '='
	const comment rune = '#'

	w, err := gowords.NewWordsFile(file, separator, comment)
	if err != nil {
		panic(err)
	}

	err = w.CheckError()
	if err != nil {
		panic(err)
	}

	value_1, found := w.Find("k1")
	if found {
		fmt.Println(value_1)
	}

	var value_2 string = w.Get("k1")
	fmt.Println(value_2)

	//Output:
	// v1
	// v1
}
