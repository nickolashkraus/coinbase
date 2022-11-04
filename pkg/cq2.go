package pkg

// Write a function that takes two users' browsing histories as input and
// returns the longest contiguous sequence of URLs that appears in both.

// This is a "longest common substring" problem and can be solved using dynamic
// programming.
//
// Algorithm:
//
// Given array A and array B, we deduce that, if a common subarray of A and B
// exists, it must start at some A[i] and B[j].
//
// We let dp[i][j] be the longest common prefix of A[i:] and B[j:]. For
// example, given the following arrays:
//
//	arr1 = [1, 2, 3, 4]
//	arr2 = [4, 2, 3, 1]
//
//	dp =
//	        0  0  0  1
//	        0  1  2  0
//	        0  0  1  0
//	        1  0  0  0
//
// Whenever A[i] == B[j], we know dp[i][j] = dp[i-1][j-1] + 1.
//
// The maximum subarray is then max(dp[i][j]) over all i, j.
//
// The time complexity of this solution is O(n*m).
func findContiguousHistory(u1 []string, u2 []string) []string {
	// Allocate a zeroed 2D array of len(u1) x len(u2).
	dp := make([][]int, len(u1))
	for i := 0; i < len(u1); i++ {
		dp[i] = make([]int, len(u2))
	}
	max, start := 0, 0
	for i := 0; i < len(u1); i++ {
		for j := 0; j < len(u2); j++ {
			if u1[i] == u2[j] {
				// Handle case where the first element of either array begins a common
				// substring. When taking a bottom-up approach, this check is
				// unneccessary.
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > max {
					max = dp[i][j]
					start = i - (max - 1)
				}
			}
		}
	}
	// If max != zero, simply return a slice of u1 using the starting index and
	// the longest common prefix (int) taken from dp.
	if max == 0 {
		return make([]string, 0)
	} else {
		return u1[start : start+max]
	}
}
