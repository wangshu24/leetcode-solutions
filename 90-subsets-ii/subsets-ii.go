func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    slices.Sort(nums)

    var backtrack func(i int, subset []int)
    backtrack = func(i int, subset []int) {
        if i == len(nums) {
            dst := make([]int, len(subset))
            copy(dst, subset)
            res = append(res, dst)
            return
        } else if i > len(nums) {return}

        subset = append(subset, nums[i])
        backtrack(i+1, subset)
        subset = subset[: len(subset)-1]
        for i+1 < len(nums) && nums[i] == nums[i+1] {
            i++
        }
        backtrack(i+1, subset)
    }

    backtrack(0, make([]int, 0))
    return res
}