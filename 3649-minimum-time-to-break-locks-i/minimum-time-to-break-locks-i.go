import (
	"slices"
)

func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	slices.Sort(strength)

	ans := 1e9

	var dfs func(mask, power, time int)
	dfs = func(mask, power, time int) {
		if mask == (1<<n)-1 {
			if float64(time) < ans {
				ans = float64(time)
			}
			return
		}

		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				continue
			}
			add := (strength[i] + power - 1) / power
			dfs(mask|(1<<i), power+k, time+add)
		}
	}

	dfs(0, 1, 0)
	return int(ans)
}
