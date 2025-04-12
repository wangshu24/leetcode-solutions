func subsetXORSum(nums []int) int {
    var res int

    var backtrack func(i, cursum int)
    backtrack = func(i, cursum int) {
        if i == len(nums) {
            res += cursum
            return
        } else if i > len(nums) {return}
        prev := cursum
        cursum = cursum ^ nums[i]
        backtrack(i+1, cursum)
        backtrack(i+1, prev)
    }

    backtrack(0, 0)

    return res
}