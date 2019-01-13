# Algorithms in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ilango100/algo)](https://goreportcard.com/report/github.com/ilango100/algo)
[![GoDoc](https://godoc.org/github.com/ilango100/algo?status.svg)](https://godoc.org/github.com/ilango100/algo)

[![Build Status](https://travis-ci.org/ilango100/algo.svg?branch=master)](https://travis-ci.org/ilango100/algo)
[![Coverage Status](https://coveralls.io/repos/github/ilango100/algo/badge.svg?branch=master)](https://coveralls.io/github/ilango100/algo?branch=master)

This repo provides implementation of basic algorithms in go.

## Algorithms implemented

### Sorting
- Quicksort - [Wiki](https://en.wikipedia.org/wiki/Quicksort)
- Selection Sort - [Wiki](https://en.wikipedia.org/wiki/Selection_sort)
- Insertion Sort - [Wiki](https://en.wikipedia.org/wiki/Insertion_sort)

### Maximum Subarray - [Wiki](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
- Brute Force
- Divide, Conquer & Combine (DC)
- Kadane's Algorithm

```
$ go test -benchmem -bench . ./...
goos: windows
goarch: amd64
pkg: github.com/ilango100/algo/sort
BenchmarkSortInsertion-4            5000            329518 ns/op               0 B/op          0 allocs/op
BenchmarkSortQuick-4              100000             19049 ns/op               0 B/op          0 allocs/op
BenchmarkSortSelection-4            3000            483373 ns/op               0 B/op          0 allocs/op
BenchmarkSortGo-4                  20000             66123 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/ilango100/algo/sort  7.624s
goos: windows
goarch: amd64
pkg: github.com/ilango100/algo/subarray
BenchmarkSubarrayBruteForce-4                 20          73403570 ns/op               0 B/op          0 allocs/op
BenchmarkSubarrayDC-4                     100000             16256 ns/op               0 B/op          0 allocs/op
BenchmarkSubarrayKadane-4                1000000              1371 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/ilango100/algo/subarray      5.162s
```

This is a **Work in Progress**
