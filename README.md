go-words
========

[![Release](https://img.shields.io/github/v/release/saleh-rahimzadeh/go-words?filter=v1.2.0&label=Release)](https://github.com/saleh-rahimzadeh/go-words/releases/tag/v1.2.0)
![Go version](https://img.shields.io/github/go-mod/go-version/saleh-rahimzadeh/go-words)
[![Go Reference](https://pkg.go.dev/badge/github.com/saleh-rahimzadeh/go-words.svg)](https://pkg.go.dev/github.com/saleh-rahimzadeh/go-words)
[![Test Status](https://github.com/saleh-rahimzadeh/go-words/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/saleh-rahimzadeh/go-words/actions/workflows/test.yml?query=branch%3Amain)
![Coverage](https://img.shields.io/badge/Coverage-95.7%25-brightgreen)
[![codecov](https://codecov.io/gh/saleh-rahimzadeh/go-words/graph/badge.svg?token=O4EXLIR5ZN)](https://codecov.io/gh/saleh-rahimzadeh/go-words)
[![Go Report Card](https://goreportcard.com/badge/github.com/saleh-rahimzadeh/go-words)](https://goreportcard.com/report/github.com/saleh-rahimzadeh/go-words)
<!-- [![Awesome Go](https://awesome.re/mentioned-badge.svg)](https://awesome-go.com/translation) -->

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

A source is a string or a file that contains text lines, each separated by a line break.

Each line must contain a _key_ and _value_ pair, with the key separated from the value by a separator character.

A **key** must be a unique and searchable word.
It can be a single word or a phrase.

A **value** can be a single word, a phrase, or empty.

All leading and trailing whitespace of the key and value will be trimmed and removed upon loading.

A **separator** is a single delimiter character that separates the key and value in each line.
The default separator character is `=` (of `rune` type).
Other valid characters such as `|`, `:`, `;`, `,`, `.`, `?`, and `@` can also be used.

A line can also be a comment or an empty line.

A **comment** is a line that starts with a comment character.
The default comment character is `#` (of `rune` type), but other valid characters such as `|`, `:`, `;`, `,`, `.`, `?`, and `@` can also be used.

All comments and empty lines will be removed upon loading.

You cannot use the same character for both the separator and the comment.

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
// Source string
const stringSource string = `
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
```



## APIs

The **go-words** contain 3 different types of APIs, each with a different source and storage, so they have different performances, throughput, and resource usage.

| API               | Source     | Storage    | Source Validation    | Resource Usage |
|-------------------|------------|------------|----------------------|----------------|
| `WordsRepository` | `string`   | array      | On instantiation     | Memory         |
| `WordsCollection` | `string`   | map        | On instantiation     | Memory         |
| `WordsFile`       | `*os.File` | `*os.File` | Calling `CheckError` | CPU            |



## Instantiation

To create `WordsRepository` and `WordsCollection` instances use `NewWordsRepository`, `NewWordsCollection` functions and provide source, separator character and comment character.

`WordsRepository` example:

```go
func main() {
  const separator = '='
  const comment = '#'
  var err error

  var wrd gowords.WordsRepository
  wrd, err = gowords.NewWordsRepository(stringSource, separator, comment)
}
```

`WordsCollection` example:

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

2. Define a variable of `Suffix` type from `core` package and instantiate `WithSuffix` API with a instance of `Words` interface (an instance of `NewWordsRepository`, `NewWordsCollection`, and `NewWordsFile`) and defined suffix.

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
  println(value2en, found2en)  // OUTPUT: "Value 2 English", true

  value1fa := wordsFA.Get("key1")
  println(value1fa)  // OUTPUT: "Value 1 Farsi"

  value2fa, found2fa := wordsFA.Find("key2")
  println(value2fa, found2fa)  // OUTPUT: "Value 2 Farsi", true
}
```

The `NewWithSuffix` function validate suffix on calling.



## Annotations

Using `DoAnnotation` API to format value according to an annotation or a format specifier.

There are 3 types of annotations:
- Named: format value using named tokens like `{{name}}` by `GetNamed` and `FindNamed` methods.
- Indexed: format value using indexed tokens like `{{1}}` by `GetIndexed` and `FindIndexed` methods.
- Formatted: format value using formatted verbs (https://pkg.go.dev/fmt#hdr-Printing) like `%s` by `GetFormatted` and `FindFormatted` methods.

```go
func main() {
  const stringSource string = `
key_named = Application {{name}} , Version {{ver}}.
key_indexed = Application {{1}} , Version {{2}}.
key_formatted = Application %s , Version %d.
`

  var wRepository gowords.WordsRepository
  wRepository, err := gowords.NewWordsRepository(stringSource, core.Separator, core.Comment)
  if err != nil {
    panic(err)
  }

  words, err := gowords.NewDoAnnotation(wRepository)
  if err != nil {
    panic(err)
  }

  value1 := words.GetNamed("key_named", map[string]string{
    "name": "MyAppX",
    "age": "111",
  })
  fmt.Println(value1)  // OUTPUT: "Application MyAppX , Version 111"

  value2, found2 := words.FindNamed("key_named", map[string]string{
    "name": "MyAppZ",
    "age": "222",
  })
  fmt.Println(value2, found2)  // OUTPUT: "Application MyAppZ , Version 222", true

  value3 := words.GetIndexed("key_indexed", "MyAppQ", 333)
  fmt.Println(value3)  // OUTPUT: "Application MyAppQ , Version 333"

  value4, found4 := words.FindIndexed("key_indexed", "MyAppW", 444)
  fmt.Println(value4, found4)  // OUTPUT: "Application MyAppW , Version 444"

  value5 := words.GetFormatted("key_formatted", "MyAppN", 555)
  fmt.Println(value5)  // OUTPUT: "Application MyAppN , Version 555"

  value6, found6 := words.FindFormatted("key_formatted", "MyAppM", 666)
  fmt.Println(value6, found6)  // OUTPUT: "Application MyAppM , Version 666"
}
```



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
BenchmarkWordsCollection-4            20426679        62.60 ns/op          0 B/op       0 allocs/op
BenchmarkWordsRepository-4               65107        16248 ns/op          1 B/op       0 allocs/op
BenchmarkWordsFile-4                      9357       191994 ns/op      19280 B/op    1001 allocs/op
BenchmarkWordsFileUnsafe-4                5299       238233 ns/op      19280 B/op    1001 allocs/op
BenchmarkDoAnnotationNamed-4            352963         4641 ns/op        152 B/op       6 allocs/op
BenchmarkDoAnnotationIndexed-4          171255         6257 ns/op        498 B/op       9 allocs/op
BenchmarkDoAnnotationFormatted-4       1000000         1040 ns/op         64 B/op       1 allocs/op
PASS
coverage: 46.2% of statements
ok    github.com/saleh-rahimzadeh/go-words  6.400s
```



## Architecture Decisions

Architecture decision records (ADR) and design specifications:

| Index                        | Description                    |
| ---------------------------- | ------------------------------ |
| [01](doc/architecture-01.md) | Deciding on a parsing strategy |
| [02](doc/architecture-02.md) | Providing different storages   |
