# 01. Deciding on a parsing strategy

| Metadata | Value             |
| -------- | ----------------- |
| Date     | 2023-12-14        |
| Author   | @saleh-rahimzadeh |
| Status   | Accepted          |

## Context

Choosing a parsing strategy, whether string operations or regex, to parse lines and key/value pairs.

## Decision

In Go, string operations are faster than regex, so we decided to parse lines and the key/value pairs using string operations. 

Consider the following code, a comparison between string operations and regex to parse lines:

```go
package main

import (
  "regexp"
  "strings"
)

const newline   = "\n"
var rxNewline   = regexp.MustCompile(newline)
const find      = "name"
const separator = "="
const source    = 
`k0=v0
k1=v1
k2=v2
k3=v3
k4=v4
k5=v5
k6=v6
k7=k7
k8=k8
k9=k9
name=Go
`

func UseRegex() string {
  for _, row := range rxNewline.Split(source, -1) {
    key, value, _ := strings.Cut(row, separator)
    if key == find {
      return value
    }
  }
  panic("not found")
}

func UseString() string {
  for _, row := range strings.Split(source, newline) {
    key, value, _ := strings.Cut(row, separator)
    if key == find {
      return value
    }
  }
  panic("not found")
}
```

Benchmark result:

```txt
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i3 CPU     @ 2.93GHz
BenchmarkUseRegex-4       253411     7100 ns/op     1428 B/op     15 allocs/op
BenchmarkUseString-4     1163377     1119 ns/op      192 B/op      1 allocs/op
PASS
ok    3.689s
```

## Consequences

Using string operations is more than 6x faster than regex and uses less memory.
