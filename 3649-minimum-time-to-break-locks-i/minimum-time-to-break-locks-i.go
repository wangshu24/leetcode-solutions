func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	tot := 1 << n
	dp := make([]int, tot)
	for i := 0; i < tot; i++ {
		dp[i] = -1
	}
	var rec func(mask int) int
	rec = func(mask int) int {
		if mask == tot-1 {
			return 0
		}
		if dp[mask] != -1 {
			return dp[mask]
		}
		cnt := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				cnt++
			}
		}
		v := 1 + cnt*k
		best := 1 << 30
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				t := (strength[i] + v - 1) / v
				cur := t + rec(mask|(1<<i))
				if cur < best {
					best = cur
				}
			}
		}
		dp[mask] = best
		return best
	}
	return rec(0)
}