func maxAbsoluteSum(nums []int) int {
    maxV, minV := nums[0], nums[0]
    maxEnd, minEnd := nums[0], nums[0]
    var res int

    for i:=1; i < len(nums); i++ {
        maxEnd = max(maxEnd+nums[i], nums[i])
        maxV = max(maxEnd, maxV)
    }

    for i:=1; i < len(nums); i++{
        minEnd = min(minEnd + nums[i], nums[i])
        minV = min(minV, minEnd)
    }

    res = max(maxV, -minV)
    return res
}

func max(a, b int) int { 
    if a  > b {return a}
    return b
}
func min(a,b  int) int {
    if a >b {return b}
    return a
}