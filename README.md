# Algorithms in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ilango100/algo)](https://goreportcard.com/report/github.com/ilango100/algo)
[![GoDoc](https://godoc.org/github.com/ilango100/algo?status.svg)](https://godoc.org/github.com/ilango100/algo)

This repo provides implementation of basic algorithms in go.

## Algorithms implemented

### Sorting
- Quicksort - [Wiki](https://en.wikipedia.org/wiki/Quicksort)
- Selection Sort - [Wiki](https://en.wikipedia.org/wiki/Selection_sort)
- Insertion Sort - [Wiki](https://en.wikipedia.org/wiki/Insertion_sort)

```
$ go test -benchmem -bench . ./...
goos: windows
goarch: amd64
pkg: github.com/ilango100/algo/sort
BenchmarkSortInsertion-4            2000            917046 ns/op               0 B/op          0 allocs/op
BenchmarkSortQuick-4               30000             46477 ns/op               0 B/op          0 allocs/op
BenchmarkSortSelection-4            1000           1045234 ns/op               0 B/op          0 allocs/op
BenchmarkSortGo-4                  10000            145211 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/ilango100/algo/sort  7.232s
```

This is a **Work in Progress**
