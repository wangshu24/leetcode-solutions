func permute(nums []int) [][]int {
    var res [][]int
    visit := map[int]bool{}

    var backtrack func(i int, subset []int, path map[int]bool)
    backtrack = func(i int, subset []int, path map[int]bool) {
        if len(subset) == len(nums) {
            dst := make([]int, len(nums))
            copy(dst, subset)
            res = append(res, dst)
            return
        } else if i >= len(nums) || len(subset) > len(nums) {
            return
        }

        for _,j := range nums {
            if !path[j] {
                subset = append(subset, j)
                path[j] = true
                backtrack(i+1, subset, path)
                subset = subset[:len(subset)-1]
                path[j] = false
            }
        }
    }

    backtrack(0, make([]int, 0), visit)
    
    return res
}