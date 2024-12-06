func maxCount(banned []int, n int, maxSum int) int {
    sort.Ints(banned)
    sum := 0
    res := 0
    for i := 1; i <= n; i++ {
        if len(banned) > 0 && i == banned[0] {
            for len(banned) > 0 && i == banned[0] {
                banned = banned[1: ]
            }
        } else {
            sum += i
            res++
        }
        if sum > maxSum {
            return res - 1
        }
    }
    return res
}