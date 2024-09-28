func searchInsert(nums []int, target int) int {
    newNums := append(nums, target)
    slices.Sort(newNums)
    result := slices.IndexFunc(newNums, func(n int) bool {
		return n == target
	})
    return result
}