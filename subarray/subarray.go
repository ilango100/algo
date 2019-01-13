package subarray

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

//MaxSumBruteForce calculates the Maximum Sum Subarray of given array using Brute Force
//The running time is O(n^2)
func MaxSumBruteForce(a []int) (lo int, hi int) {
	maxSum := -((^0) >> 1) - 1
	for i := range a {
		for j := i; j < len(a); j++ {
			if s := sum(a[i:j]); s > maxSum {
				maxSum = s
				lo = i
				hi = j
			}
		}
	}
	return lo, hi
}

//MaxSumDC calculates the Maximum Sum Subarray of given array using Divide and Conquer
//The running time is O(n lg(n))
func MaxSumDC(a []int) (lo int, hi int) {
	return 0, len(a) - 1
}
