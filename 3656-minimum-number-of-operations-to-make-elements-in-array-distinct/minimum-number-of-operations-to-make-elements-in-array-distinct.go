func minimumOperations(nums []int) int {
    s, l := slices.Min(nums), slices.Max(nums)
    tmp := make([]int, l-s+1)
    for _,i := range nums {
        tmp[i-s]++
    }
    var min int
    if slices.Max(tmp) == 1 {return min}

    for i:=0; i < len(nums); i+=3{
        if i >= len(nums) || i+2 >= len(nums) {break}
        tmp[nums[i]-s]--
        tmp[nums[i+1]-s]--
        tmp[nums[i+2]-s]--
        min++
        if slices.Max(tmp) == 1 {return min}
    }
    
    if slices.Max(tmp) > 1 {return min+1}
    return min
}