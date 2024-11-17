func shortestSubarray(nums []int, k int) int {
    n := len(nums)
    // Create prefix sum array
    prefix := make([]int64, n+1)
    for i := 0; i < n; i++ {
        prefix[i+1] = prefix[i] + int64(nums[i])
    }
    
    // Initialize result and deque
    result := n + 1
    deque := make([]int, 0, n+1)
    
    for i := 0; i <= n; i++ {
        // Remove indices that break monotonicity
        for len(deque) > 0 && prefix[i] <= prefix[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        
        // Check for valid subarrays
        for len(deque) > 0 && prefix[i]-prefix[deque[0]] >= int64(k) {
            result = min(result, i-deque[0])
            deque = deque[1:]
        }
        
        deque = append(deque, i)
    }
    
    if result <= n {
        return result
    }
    return -1
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}