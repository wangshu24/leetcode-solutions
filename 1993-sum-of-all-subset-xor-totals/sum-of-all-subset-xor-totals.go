func subsetXORSum(nums []int) int {
    var res int

    var backtrack func(i, cur int)
    backtrack = func(i, cur int) {
        if i >= len(nums) {
            res += cur
            return
        }

        backtrack(i+1, cur)
        backtrack(i+1, cur ^ nums[i])
    } 

    backtrack(0, 0)
    return res
}