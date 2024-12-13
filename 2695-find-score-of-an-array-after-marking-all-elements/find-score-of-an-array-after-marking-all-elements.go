func findScore(nums []int) int64 {
	result := int64(0)
	start := 0 // Starting index of current decision chain
	for start < len(nums) {
		end := start
		for end < len(nums)-1 && nums[end] > nums[end+1] {
			end++
		}

		for i := end; start <= i; i -= 2 {
			result += int64(nums[i])
		}
		start = end + 2
	}
	return result
}