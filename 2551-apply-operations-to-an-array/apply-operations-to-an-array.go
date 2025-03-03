func applyOperations(nums []int) []int {
    n := len(nums)
    
    // Apply operations
    for i := 0; i < n-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        }
    }
    
    // Shift zeros to the end (in-place)
    nonZeroIdx := 0
    
    // Move all non-zero elements to the front
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            nums[nonZeroIdx] = nums[i]
            nonZeroIdx++
        }
    }
    
    // Fill the remaining positions with zeros
    for i := nonZeroIdx; i < n; i++ {
        nums[i] = 0
    }
    
    return nums
}