// countSetBits returns the number of set bits in the binary representation of a number.
func countSetBits(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

// canSortArray determines if the array can be sorted by only swapping adjacent elements
// that have the same number of set bits.
func canSortArray(nums []int) bool {
	n := len(nums)
	count := make([]int, n)

	// Calculate the number of set bits for each number in nums.
	for i, num := range nums {
		count[i] = countSetBits(num)
	}

	prevMax := -1 << 31 // Initialize to a very small integer
	for i := 0; i < n; {
		end := i
		currMax := -1 << 31
		currMin := 1 << 31 - 1 // Initialize to a large integer

		// Find the range of elements with the same number of set bits
		for j := i; j < n && count[j] == count[i]; j++ {
			if nums[j] > currMax {
				currMax = nums[j]
			}
			if nums[j] < currMin {
				currMin = nums[j]
			}
			end++
		}
		i = end

		// If the minimum of this segment is less than the previous segment's max, return false
		if currMin < prevMax {
			return false
		}
		prevMax = currMax
	}
	return true
}