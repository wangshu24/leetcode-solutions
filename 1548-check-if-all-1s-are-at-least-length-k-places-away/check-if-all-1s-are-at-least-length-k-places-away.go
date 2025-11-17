func kLengthApart(nums []int, k int) bool {
    last := slices.Index(nums, 1)
    for i:= last+1; i < len(nums); i++ {
        if nums[i] != 1  {
            continue    
        } else {
            if i-last-1 >= k {
                last=i
                fmt.Println(i)
            } else {
                return false
            }
        }
    }

    return true
}