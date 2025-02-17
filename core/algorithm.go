package core

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
