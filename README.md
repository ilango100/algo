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
BenchmarkInsertionSort-4            3000            362124 ns/op               0 B/op          0 allocs/op
BenchmarkQuickSort-4              100000             20965 ns/op               0 B/op          0 allocs/op
BenchmarkGoSort-4                  20000             72293 ns/op              32 B/op          1 allocs/op
BenchmarkSelectionSort-4            3000            518684 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/ilango100/algo/sort  7.585s
```

This is a **Work in Progress**
