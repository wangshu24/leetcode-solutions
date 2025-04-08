func combinationSum2(cand []int, target int) [][]int {
    slices.Sort(cand)
    var res [][]int

    var backtrack func(i, total int, subset []int)
    backtrack = func(i, total int,  subset []int) {
        if total == target {
            dst := make([]int, len(subset))
            copy(dst, subset)
            res = append(res, dst)
            return
        } else if i >= len(cand) || total > target {
            return
        }

        subset = append(subset, cand[i])
        backtrack(i+1, total + cand[i], subset)
        for i+1 < len(cand) && cand[i] == cand[i+1] {
            i++
        }
        subset = subset[:len(subset)-1]
        backtrack(i+1, total, subset)
    }

    backtrack(0,0, make([]int, 0))
    
    return res
}