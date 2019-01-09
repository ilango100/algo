# Algorithms in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ilango100/algo)](https://goreportcard.com/report/github.com/ilango100/algo)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ilango100/algo)

This repo provides implementation of basic algorithms in go.

## Algorithms implemented

### Sorting
- Quicksort - [Wiki](https://en.wikipedia.org/wiki/Quicksort)

```
$ go test -bench . -benchmem -args -n=10000
goos: windows
goarch: amd64
pkg: github.com/ilango100/algo
BenchmarkQSort-4                    5000            240362 ns/op               0 B/op          0 allocs/op
BenchmarkGoSort-4                   2000            871169 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/ilango100/algo       4.115s
```

This is a **Work in Progress**
