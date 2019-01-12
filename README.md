# Algorithms in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ilango100/algo)](https://goreportcard.com/report/github.com/ilango100/algo)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ilango100/algo)

This repo provides implementation of basic algorithms in go.

## Algorithms implemented

### Sorting
- Quicksort - [Wiki](https://en.wikipedia.org/wiki/Quicksort)

```
$ go test -benchmem -bench . ./... -args n=10000
goos: windows
goarch: amd64
pkg: github.com/ilango100/algo/sort
BenchmarkQuickSort-4               20000             56947 ns/op               0 B/op          0 allocs/op
BenchmarkGoSort-4                  10000            163263 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/ilango100/algo/sort  4.141s

```

This is a **Work in Progress**
