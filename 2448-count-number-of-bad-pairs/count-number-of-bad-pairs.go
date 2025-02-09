func countBadPairs(nums []int) int64 {
    good := make(map[int]int)
    res,gp := 0, 0

    for i:=0; i < len(nums); i++ {
        good[nums[i]-i]++
    }
    for _, val := range good {
        if val > 1 {
            gp += (val * (val-1))/2
        }
    }

    res = ((len(nums) * (len(nums)-1))/2)-gp
    return int64(res)
}