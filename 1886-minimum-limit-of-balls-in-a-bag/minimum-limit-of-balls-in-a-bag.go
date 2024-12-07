func minimumSize(nums []int, maxOp int) int {
    l,r  := 1, slices.Max(nums)
    
    for l < r  {
        mid := (l+r)/2
        ops := 0
        for _, n := range nums {
            ops += (n-1)/mid
        }
        if ops  <=  maxOp {
            r = mid
        } else { l = mid + 1}
    }

    return r
}