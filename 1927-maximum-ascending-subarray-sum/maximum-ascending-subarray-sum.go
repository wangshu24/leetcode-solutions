func max(a, b int) int {
    if a < b  {return b}
    return a
}

func maxAscendingSum(nums []int) int {
    res, temp :=nums[0], nums[0]

    for i:= 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            temp+=nums[i]
        } else {
            res = max(res, temp)
            temp = nums[i]
        }
    }
    res = max(res, temp)
    return res
}