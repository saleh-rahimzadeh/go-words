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
	const source string = `
k1:v1
k2:v2
k3:v3
`

	w, err := gowords.NewWordsRepository(source, separator, core.Comment)
	if err != nil {
		panic(err)
	}

	value := w.Get("k1")
	fmt.Println(value)

	//Output: v1
}

func ExampleWordsRepository_customComment() {
	const comment rune = '@'
	const source string = `
k1=v1
@ this is a comment
k2=v2
`

	w, err := gowords.NewWordsRepository(source, core.Separator, comment)
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
k1_EN = v1 EN
k2_EN = v2 EN
k1_FA = v1 FA
k2_FA = v2 FA
`

	var wRepository gowords.Words
	wRepository, err := gowords.NewWordsRepository(source, core.Separator, core.Comment)
	if err != nil {
		panic(err)
	}

	const (
		EN core.Suffix = "_EN"
		FA core.Suffix = "_FA"
	)

	wordsEN, err := gowords.NewWithSuffix(wRepository, EN)
	if err != nil {
		panic(err)
	}

	wordsFA, err := gowords.NewWithSuffix(wRepository, FA)
	if err != nil {
		panic(err)
	}

	value1en := wordsEN.Get("k1")
	fmt.Println(value1en)

	value2en, found2en := wordsEN.Find("k2")
	if found2en {
		fmt.Println(value2en)
	}

	value1fa := wordsFA.Get("k1")
	fmt.Println(value1fa)

	value2fa, found2fa := wordsFA.Find("k2")
	if found2fa {
		fmt.Println(value2fa)
	}

	//Output:
	// v1 EN
	// v2 EN
	// v1 FA
	// v2 FA
}

//┌ DoAnnotation
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleDoAnnotation() {
	const source = `
person_named=Hi, my name is {{name}}, when I was {{age}} years old I was a {{language}} developer.
person_indexed=Hi, my name is {{1}}, when I was {{2}} years old I was a {{3}} developer.
person_formatted=Hi, my name is %s, when I was %d years old I was a %v developer.
`

	var wRepository gowords.Words
	wRepository, err := gowords.NewWordsRepository(string(source), core.Separator, core.Comment)
	if err != nil {
		panic(err)
	}

	words, err := gowords.NewDoAnnotation(wRepository)
	if err != nil {
		panic(err)
	}

	value1person := words.GetNamed("person_named", map[string]interface{}{
		"name": "Saleh",
		"age": "15",
		"language": "Assembly",
	})
	fmt.Println(value1person)

	value2person, found_person := words.FindNamed("person_named", map[string]interface{}{
		"name": "Saleh",
		"age": "17",
		"language": "Pascal",
	})
	if found_person {
		fmt.Println(value2person)
	}

	value1personindexed := words.GetIndexed("person_indexed", "Saleh", "19", "C++")
	fmt.Println(value1personindexed)

	value2personindexed, found_personindexed := words.FindIndexed("person_indexed", "Saleh", "23", "CSharp")
	if found_personindexed {
		fmt.Println(value2personindexed)
	}

	value1personformatted := words.GetFormatted("person_formatted", "Saleh", 32, "JavaScript")
	fmt.Println(value1personformatted)

	value2personformatted, found_personformatted := words.FindFormatted("person_formatted", "Saleh", 36, "Golang")
	if found_personformatted {
		fmt.Println(value2personformatted)
	}

	//Output:
	// Hi, my name is Saleh, when I was 15 years old I was a Assembly developer.
	// Hi, my name is Saleh, when I was 17 years old I was a Pascal developer.
	// Hi, my name is Saleh, when I was 19 years old I was a C++ developer.
	// Hi, my name is Saleh, when I was 23 years old I was a CSharp developer.
	// Hi, my name is Saleh, when I was 32 years old I was a JavaScript developer.
	// Hi, my name is Saleh, when I was 36 years old I was a Golang developer.
}

//┌ Services Example
//└─────────────────────────────────────────────────────────────────────────────────────────────────

func ExampleGetBy() {
	const source = `
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
