import (
	"math"
)

func splitArray(nums []int, k int) int {
	n := len(nums)
	// Memoization table: key is [2]int{index, remainingK}
	memo := make(map[[2]int]int)

	var dfs func(idx int, cuts int) int
	dfs = func(idx int, cuts int) int {
		// Base Case 1: If only one partition left, the result is the sum of the rest
		if cuts == 1 {
			sum := 0
			for i := idx; i < n; i++ {
				sum += nums[i]
			}
			return sum
		}

		// Check Memo
		state := [2]int{idx, cuts}
		if val, ok := memo[state]; ok {
			return val
		}

		minLargestSum := math.MaxInt32
		currentSum := 0

		// Base Case 2: Optimization - ensure enough elements left for remaining cuts
		// We can only iterate until we have (cuts-1) elements remaining for the future splits
		for i := idx; i <= n-cuts; i++ {
			currentSum += nums[i]
			
			// Get max sum of the remaining splits
			remainingMax := dfs(i+1, cuts-1)
			
			// We want the maximum of (current_subarray_sum, result_of_sub_problem)
			// Then we want to minimize this value across all possible split points
			maxVal := max(currentSum, remainingMax)
			minLargestSum = min(minLargestSum, maxVal)
		}

		memo[state] = minLargestSum
		return minLargestSum
	}

	return dfs(0, k)
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a > b { return a }
	return b
}