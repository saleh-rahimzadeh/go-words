package gowords_test

import (
	"fmt"
	"os"
	"path"

	gowords "github.com/saleh-rahimzadeh/go-words"
	"github.com/saleh-rahimzadeh/go-words/core"
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

//┌ WithSuffix Examples
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleWithSuffix() {
	const source = `
k1 = v1
k1_EN = v1 EN
k2_EN = v2 EN
k1_FA = v1 FA
k2_FA = v2 FA
`

	var words gowords.Words
	words, err := gowords.NewWordsRepository(source, core.Separator, core.Comment)
	if err != nil {
		panic(err)
	}

	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)

	wEN, err := gowords.NewWithSuffix(words, EN)
	if err != nil {
		panic(err)
	}

	wFA, err := gowords.NewWithSuffix(words, FA)
	if err != nil {
		panic(err)
	}

	value1en := wEN.Get("k1")
	fmt.Println(value1en)

	value2en, found2en := wEN.Find("k2")
	if found2en {
		fmt.Println(value2en)
	}

	value1fa := wFA.Get("k1")
	fmt.Println(value1fa)

	value2fa, found2fa := wFA.Find("k2")
	if found2fa {
		fmt.Println(value2fa)
	}

	//Output:
	// v1 EN
	// v2 EN
	// v1 FA
	// v2 FA
}

//┌ Services Example
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleGetBy() {
	const source = `
k1 = v1
k1_EN = v1 EN
k2_EN = v2 EN
k1_FA = v1 FA
k2_FA = v2 FA
`

	var words gowords.Words
	words, err := gowords.NewWordsRepository(source, core.Separator, core.Comment)
	if err != nil {
		panic(err)
	}

	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)

	valueEn := gowords.GetBy(words, "k1", EN)
	fmt.Println(valueEn)

	valueFa := gowords.GetBy(words, "k1", FA)
	fmt.Println(valueFa)

	//Output:
	// v1 EN
	// v1 FA
}

func ExampleFindBy() {
	const source = `
k1 = v1
k1_EN = v1 EN
k2_EN = v2 EN
k1_FA = v1 FA
k2_FA = v2 FA
`

	var words gowords.Words
	words, err := gowords.NewWordsRepository(source, core.Separator, core.Comment)
	if err != nil {
		panic(err)
	}

	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)

	valueEn, foundEn := gowords.FindBy(words, "k1", EN)
	if foundEn {
		fmt.Println(valueEn)
	}

	valueFa, foundFa := gowords.FindBy(words, "k1", FA)
	if foundFa {
		fmt.Println(valueFa)
	}

	//Output:
	// v1 EN
	// v1 FA
}
