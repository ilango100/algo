# Algorithms in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ilango100/algo)](https://goreportcard.com/report/github.com/ilango100/algo)
[![GoDoc](https://godoc.org/github.com/ilango100/algo?status.svg)](https://godoc.org/github.com/ilango100/algo)

[![Build Status](https://travis-ci.org/ilango100/algo.svg?branch=master)](https://travis-ci.org/ilango100/algo)
[![Coverage Status](https://coveralls.io/repos/github/ilango100/algo/badge.svg?branch=master)](https://coveralls.io/github/ilango100/algo?branch=master)

This repo provides implementation of basic algorithms in go.

## Algorithms implemented

### Sorting
- Quicksort - [Wiki](https://en.wikipedia.org/wiki/Quicksort)
- Heapsort - [Wiki](https://en.wikipedia.org/wiki/Heapsort)
- Selection Sort - [Wiki](https://en.wikipedia.org/wiki/Selection_sort)
- Insertion Sort - [Wiki](https://en.wikipedia.org/wiki/Insertion_sort)
```
BenchmarkSortHeap-4                20000             68416 ns/op               0 B/op          0 allocs/op
BenchmarkSortInsertion-4            5000            254319 ns/op               0 B/op          0 allocs/op
BenchmarkSortQuick-4              100000             19747 ns/op               0 B/op          0 allocs/op
BenchmarkSortSelection-4            3000            489689 ns/op               0 B/op          0 allocs/op
BenchmarkSortGo-4                  20000             66971 ns/op              32 B/op          1 allocs/op
```

### Maximum Subarray - [Wiki](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
- Brute Force
- Divide, Conquer & Combine (DC)
- Kadane's Algorithm
```
BenchmarkSubarrayBruteForce-4                 20          74001950 ns/op               0 B/op          0 allocs/op
BenchmarkSubarrayDC-4                     100000             16306 ns/op               0 B/op          0 allocs/op
BenchmarkSubarrayKadane-4                1000000              1371 ns/op               0 B/op          0 allocs/op
```

## Datastructures

- Queue - [Wiki](https://en.wikipedia.org/wiki/Queue_(abstract_data_type))
- Stack - [Wiki](https://en.wikipedia.org/wiki/Stack_(abstract_data_type))
- Binary Heap - [Wiki](https://en.wikipedia.org/wiki/Binary_heap)

This is a **Work in Progress**
