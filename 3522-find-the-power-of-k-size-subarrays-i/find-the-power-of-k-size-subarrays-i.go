func resultsArray(nums []int, k int) []int {
    if k == 1 {
        return nums
    }
    
    n := len(nums)
    result := make([]int, 0)
    left := 0
    right := 1
    
    for right < n {
        isNotConsecutive := nums[right] - nums[right-1] != 1
        
        if isNotConsecutive {
            for left < right && left + k - 1 < n {
                result = append(result, -1)
                left++
            }
            left = right
        } else if right - left == k - 1 {
            result = append(result, nums[right])
            left++
        }
        
        right++
    }
    
    return result
}