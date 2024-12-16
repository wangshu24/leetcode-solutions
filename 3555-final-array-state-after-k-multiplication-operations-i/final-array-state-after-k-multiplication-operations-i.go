import ("slices")

func getFinalState(nums []int, k int, mul int) []int {
    for k > 0 {
        min := slices.Min(nums)
        ind := slices.Index(nums, min)
        nums[ind] = min * mul
        k--
    }

    return nums
}