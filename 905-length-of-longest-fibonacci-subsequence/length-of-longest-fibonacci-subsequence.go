func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, n)
	}

	maxLen := 0
	for curr := 2; curr < n; curr++ {
		start := 0
		end := curr - 1
		finalSum := arr[curr]
		for start < end {
			currSum := arr[start] + arr[end]
			if currSum > finalSum {
				end--
			} else if currSum < finalSum {
				start++
			} else {
				dp[end][curr] = dp[start][end] + 1
				maxLen = max(maxLen, dp[end][curr])
				end--
				start++
			}
		}
	}
	if maxLen == 0 {
		return 0
	}
	return maxLen + 2
}