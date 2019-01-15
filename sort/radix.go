package sort

import (
	"math"
)

func RadixSort(a []int) {
	d := 0
	for _, v := range a {
		if v > d {
			d = v
		}
	}
	for i := 1; d > 0; i++ {
		radixDigit(a, i)
		d /= 10
	}
}

func radixDigit(b []int, j int) {
	a := make([]int, len(b))
	copy(a, b)
	j = int(math.Pow10(j))
	d := make([]int, len(b))
	for i, v := range b {
		d[i] = (v % j) / (j / 10)
	}
	c := make([]int, 10)
	for i := range a {
		c[d[i]]++
	}
	for i := 1; i < len(c); i++ {
		c[i] = c[i] + c[i-1]
	}
	for i := len(a) - 1; i >= 0; i-- {
		c[d[i]]--
		b[c[d[i]]] = a[i]
	}
}
