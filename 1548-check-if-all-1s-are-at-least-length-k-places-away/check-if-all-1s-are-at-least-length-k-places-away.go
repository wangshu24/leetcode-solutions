func kLengthApart(nums []int, k int) bool {
    last := slices.Index(nums, 1)
    fmt.Println(last) 
    for i:= last+1; i < len(nums); i++ {
        if nums[i] != 1  {
            continue    
        } else {
            if i-last-1 >= k {
                last=i
                fmt.Println(i)
            } else {
                fmt.Println("found 1 at index: %s", i)
                return false
            }
        }
    }

    return true
}