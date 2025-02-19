package core

// Given an integer array nums, find the subarray
// with the largest sum, and return its sum.
func MaximumSubArray(arr []int) int {
	maxSum := arr[0]
	curSum := 0
	for _, n := range arr {
		if curSum < 0 {
			curSum = 0
		}
		curSum += n
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
