package subarray

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

var ninf = -((^0) >> 1) - 1

//MaxSumBruteForce calculates the Maximum Sum Subarray of given array using Brute Force
//The running time is O(n^2)
func MaxSumBruteForce(a []int) (maxSum int, lo int, hi int) {
	maxSum = ninf
	for i := range a {
		for j := i; j < len(a); j++ {
			if s := sum(a[i:j]); s > maxSum {
				maxSum = s
				lo = i
				hi = j
			}
		}
	}
	return maxSum, lo, hi
}

//MaxSumDC calculates the Maximum Sum Subarray of given array using Divide and Conquer
//The running time is O(n lg(n))
func MaxSumDC(a []int) (maxSum int, lo int, hi int) {
	if len(a) == 1 {
		return a[0], 0, 1
	}
	mid := len(a) / 2
	ls, li, lj := MaxSumDC(a[0:mid])
	rs, ri, rj := MaxSumDC(a[mid:len(a)])
	ri, rj = mid+ri, mid+rj
	//Find cross sum
	cs, ci, cj := a[mid], mid, mid+1
	for sm, i := cs, mid-1; i >= 0; i-- {
		sm += a[i]
		if sm > cs {
			cs = sm
			ci = i
		}
	}
	for sm, j := cs, mid+1; j < len(a); j++ {
		sm += a[j]
		if sm > cs {
			cs = sm
			cj = j
		}
	}
	if ls > rs && ls > cs {
		return ls, li, lj
	} else if rs > ls && rs > cs {
		return rs, ri, rj
	}
	return cs, ci, cj
}

//MaxSumKadane calculates the maximum subarray of the given array using Kadane's Algorithm:
//It runs in O(n)
func MaxSumKadane(a []int) (maxSum int, lo int, hi int) {
	maxSum = ninf
	isum := ninf
	ilo := 0
	for i := range a {
		if a[i] > isum+a[i] {
			ilo = i
			isum = a[i]
		} else {
			isum += a[i]
		}
		if isum > maxSum {
			lo = ilo
			hi = i
			maxSum = isum
		}
	}
	return maxSum, lo, hi
}
