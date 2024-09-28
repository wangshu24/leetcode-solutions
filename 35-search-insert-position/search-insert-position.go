func searchInsert(nums []int, target int) int {
    newNums := append(nums, target)
    slices.Sort(newNums)
    result := slices.Index(newNums, target)
    return result
}