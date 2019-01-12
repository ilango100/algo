package sort

import (
	"flag"
	"math/rand"
	"os"
	"testing"
)

var a []int
var n uint

func TestMain(m *testing.M) {
	flag.UintVar(&n, "n", 1000, "Number of inputs")
	flag.Parse()
	a = rand.Perm(int(n))
	os.Exit(m.Run())
}

func testFunc(t *testing.T, f Func) {
	arr := copyArr()
	f(arr)
	testSorted(t, arr)
	testContainsAll(t, arr)
}

func benchFunc(b *testing.B, f Func) {
	arr := copyArr()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(arr)
	}

}

func testSorted(t *testing.T, arr []int) {
	e := len(a) - 2
	for i := 0; i < e; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("%d [%d] came before %d [%d]:", arr[i], i, arr[i+1], i+1)
			return
		}
	}
}

func testContainsAll(t *testing.T, arr []int) {
	for i := range arr {
		for j, v := range arr {
			if v == i {
				break
			}
			if j == len(arr)-1 {
				t.Errorf("Array did not contain %d, %v", i, arr[:j+1])
				return
			}
		}
	}
}

func copyArr() []int {
	arr := make([]int, len(a))
	copy(arr, a)
	return arr
}
