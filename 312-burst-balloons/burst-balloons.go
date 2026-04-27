func maxCoins(nums []int) int {
	temp := make([]int, 0, len(nums)+2)
	temp = append(temp, 1)
	for _, v := range nums {
		
			temp = append(temp, v)
		
	}
	temp = append(temp, 1)
	n := len(temp)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length < n; length++ {
		for i := 0; i < n-length; i++ {
			j := i + length
			for k := i + 1; k < j; k++ {
				currentBurst := temp[i] * temp[k] * temp[j]
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+currentBurst)
			}
		}
	}
	return dp[0][n-1]
}