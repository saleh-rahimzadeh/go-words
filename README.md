go-words
========

[![Release](https://img.shields.io/github/v/release/saleh-rahimzadeh/go-words)](https://github.com/saleh-rahimzadeh/go-words/releases/latest)
![Go version](https://img.shields.io/github/go-mod/go-version/saleh-rahimzadeh/go-words)
[![Go Reference](https://pkg.go.dev/badge/github.com/saleh-rahimzadeh/go-words.svg)](https://pkg.go.dev/github.com/saleh-rahimzadeh/go-words)
[![Test Status](https://github.com/saleh-rahimzadeh/go-words/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/saleh-rahimzadeh/go-words/actions/workflows/test.yml?query=branch%3Amain)
![Coverage](https://img.shields.io/badge/Coverage-94.8%25-brightgreen)
[![codecov](https://codecov.io/gh/saleh-rahimzadeh/go-words/graph/badge.svg?token=O4EXLIR5ZN)](https://codecov.io/gh/saleh-rahimzadeh/go-words)
[![Go Report Card](https://goreportcard.com/badge/github.com/saleh-rahimzadeh/go-words)](https://goreportcard.com/report/github.com/saleh-rahimzadeh/go-words)
[![Awesome Go](https://awesome.re/mentioned-badge.svg)](https://awesome-go.com/translation)

The **go-words** is a words table and text resource library for Golang projects.

It provides text for your application messages, alerts, translations, prompts and ...



## Install

To install it:

```sh
go get -u github.com/saleh-rahimzadeh/go-words
```

Import:

```go
import (
   "github.com/saleh-rahimzadeh/go-words"
   "github.com/saleh-rahimzadeh/go-words/core"
)
/* OR */
import (
   gowords "github.com/saleh-rahimzadeh/go-words"
   core    "github.com/saleh-rahimzadeh/go-words/core"
)
```



## Source

A source is a string or a file which contains text lines (each line separated by line break).

Each line must contains a _key_ and _value_ pair which the key has been separated by a separator character from value .

A **key** must be a unique word and searchable.
A key can be a single word or multiple words string.

A **value** can be a single word or multiple words string (or empty).

All leading and trailing white spaces of key and value will be trimmed and removed on loading.

A **separator** is a single delimiter character which separate key and value in each line.
Default separator character is `=` (`rune` type) and also you can use other valid characters such as `|`, `:`, `;`, `,`, `.`, `?` and `@`.

A line also can be a comment or empty line.

A **comment** is a line which started with a comment character.
Default comment character is `#` (`rune` type) and also you can use other valid characters such as `|`, `:`, `;`, `,`, `.`, `?` and `@`.

All comments and empty lines will be removed on loading.

You can't use same character for both separator and comment.

Source Outline:

```txt
<key>=<value>
# comment line
```

Key and Value samples:

| Sample                                          | Result                                |
|-------------------------------------------------|---------------------------------------|
| `key1=value1`                                   | `key1=value1`                         |
| `key2 =value2`                                  | `key2=value2`                         |
| `key3= value3`                                  | `key3=value3`                         |
| `key4 = value4`                                 | `key4=value4`                         |
| `key5    =   value5`                            | `key5=value5`                         |
| `   key6    =      value6   `                   | `key6=value6`                         |
| `key7 multi words = value7 multi words`         | `key7 multi words=value7 multi words` |
| `  key8 multi words   =   value8 multi words  ` | `key8 multi words=value8 multi words` |
| `key9=`                                         | `key9=`                               |
| `key10`                                         | ERROR separator not found             |
| `= value11`                                     | ERROR key not found                   |
| `# comment`                                     | REMOVE                                |
| `#a long long long comment`                     | REMOVE                                |
| ` `  (empty line)                               | REMOVE                                |

Source sample:

```txt
App = MyApp
Version = 1.0
Descriptin = An enterprise application
# Config Section
size = 100,200,300,400
left location = 10
right location = 20
load_position = top left
```

Suppling source in code:

```go
func main() {
   // Source string
   var stringSource string = `
App = MyApp
Version = 1.0
Descriptin = An enterprise application
# Config Section
size = 100,200,300,400
left location = 10
right location = 20
load_position = top left
`
   // Source file
   var fileSource *os.File
   fileSource, err := os.Open("<path_to_string_file>")
}
```



## APIs

GoWords contains 3 type of APIs which each one have different source and storage, so they have different performance.

| API               | Source     | Storage    | Source Validation    | Resource Usage |
|-------------------|------------|------------|----------------------|----------------|
| `WordsRepository` | `string`   | array      | On instantiation     | Memory         |
| `WordsCollection` | `string`   | map        | On instantiation     | Memory         |
| `WordsFile`       | `*os.File` | `*os.File` | Calling `CheckError` | CPU            |



## Instantiation

To create `WordsRepository` and `WordsCollection` instances use `NewWordsRepository`, `NewWordsCollection` functions and provide source, separator character and comment character.

WordsRepository example:

```go
func main() {
   const separator = '='
   const comment = '#'
   var err error

   var wrd gowords.WordsRepository
   wrd, err = gowords.NewWordsRepository(stringSource, separator, comment)
}
```

WordsCollection example:

```go
func main() {
   const separator = '='
   const comment = '#'
   var err error

   var wrd gowords.WordsCollection
   wrd, err = gowords.NewWordsCollection(stringSource, separator, comment)
}
```

These functions check and validate source, separator character, comment character and duplication on calling.

To create `WordsFile` instance use `NewWordsFile` function and provide source, separator character and comment character.

```go
func main() {
   const separator = '='
   const comment = '#'
   var err error

   var fileSource *os.File
   fileSource, err = os.Open("<path_to_string_file>")
   defer fileSource.Close()

   var wrd gowords.NewWordsFile
   wrd, err = gowords.NewWordsFile(fileSource, separator, comment)
}
```

This function check and validate separator character and comment character on calling.
To validate source and check duplication call `CheckError` method after instantiation.

```go
err := wrd.CheckError()
```

### Delimiters

You can use pre-declared characters for separator and comment delimiters of `github.com/saleh-rahimzadeh/go-words/core` package in instantiation.

- `core.Separator` for separator.
- `core.Comment` for comment.

```go
gowords.NewWordsRepository(stringSource, core.Separator, core.Comment)
gowords.NewWordsCollection(stringSource, core.Separator, core.Comment)
gowords.NewWordsFile(fileSource, core.Separator, core.Comment)
```



## Usage

All APIs implement `Words` interface.
This interface has two method `Get` and `Find` to search key and return value.

- The `Get` method search for a key then return value if found, else return empty string:

```go
var value string

value = wrd.Get("App")
println(value)  // OUTPUT: "MyApp"

value = wrd.Get("uknown_key")
println(value)  // OUTPUT: ""
```

- The `Find` method search for a key then return value and `true` boolean if found, else return empty string and `false` boolean:

```go
var value string
var found bool

value, found = wrd.Find("App")
println(value, found)  // OUTPUT: "MyApp", true

value, found = wrd.Find("uknown_key")
println(value, found)  // OUTPUT: "", false
```

Both methods validate input key on calling.

Visit "[github.com/saleh-rahimzadeh/go-words/blob/main/example_test.go](https://github.com/saleh-rahimzadeh/go-words/blob/main/example_test.go)" to see more samples.



## Suffixes

Using `WithSuffix` API to provide categorized words table and text resource, usually for internationalization and multi language texts.

To use `WithSuffix`:

1. Provide keys of source with desire suffixes.

   It's better to concate your key and suffixe with `_` character.

2. Define a variable of `Suffix` type from `core` package and instantiate `WithSuffix` API with a instance of `Words` interface (a instance of `NewWordsRepository`, `NewWordsCollection`, and `NewWordsFile`) and defined suffix.

3. Using `Get` and `Find` methods of `WithSuffix` instance to search for a name which applied suffix.

```go
func main() {
   const stringSource string = `
key1_EN = Value 1 English
key1_FA = Value 1 Farsi

key2_EN = Value 2 English
key2_FA = Value 2 Farsi

key3_EN = Value 3 English
key3_FA = Value 3 Farsi
`

   const EN core.Suffix = "_EN"
   const FA core.Suffix = "_FA"

   var wrd gowords.WordsRepository
   wrd, err := gowords.NewWordsRepository(stringSource, core.Separator, core.Comment)
   if err != nil {
      panic(err)
   }

   wordsEN, err := gowords.NewWithSuffix(wrd, EN)
	if err != nil {
		panic(err)
	}

	wordsFA, err := gowords.NewWithSuffix(wrd, FA)
	if err != nil {
		panic(err)
	}

   value1en := wordsEN.Get("key1")
	println(value1en)  // OUTPUT: "Value 1 English"

	value2en, found2en := wordsEN.Find("key2")
   println(value2en)  // OUTPUT: "Value 2 English"

	value1fa := wordsFA.Get("key1")
	println(value1fa)  // OUTPUT: "Value 1 Farsi"

	value2fa, found2fa := wordsFA.Find("key2")
   println(value2fa)  // OUTPUT: "Value 2 Farsi"
}
```

The `NewWithSuffix` function validate suffix on calling.



## Helper functions

There are some service functions, providing helper and utility functions, and also a simpler interface to working with APIs:

- `GetBy`: a helper to search for a name by suffix and using `Get` method of `Words` object.

```go
const EN core.Suffix = "_EN"
value1En := gowords.GetBy(wrd, "key1", EN)
```

- `FindBy`: a helper to search for a name by suffix and using `Find` method of `Words` object.

```go
const EN core.Suffix = "_EN"
value1En, found := gowords.FindBy(wrd, "key1", EN)
```



## Internationalization and Multi-Language

To internationalization your messages, alerts and texts, leverage `WithSuffix` API.

Prior to version 1.1.0, visit [Wiki Internationalization](https://github.com/saleh-rahimzadeh/go-words/wiki/Internationalization).



## Benchmark

```txt
goos: linux
goarch: amd64
pkg: github.com/saleh-rahimzadeh/go-words
cpu: Intel(R) Core(TM) i3 CPU  @ 2.93GHz
BenchmarkWordsCollection-4    19909888     60.87 ns/op        0 B/op       0 allocs/op
BenchmarkWordsRepository-4       63777     15852 ns/op        1 B/op       0 allocs/op
BenchmarkWordsFile-4              7652    226539 ns/op    19280 B/op    1001 allocs/op
BenchmarkWordsFileUnsafe-4        5421    228342 ns/op    19280 B/op    1001 allocs/op
PASS
coverage: 41.6% of statements
ok    github.com/saleh-rahimzadeh/go-words  6.400s
```
