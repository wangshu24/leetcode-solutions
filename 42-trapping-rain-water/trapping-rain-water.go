func trap(height []int) int {
    l,r := 0,len(height)-1
    mL, mR := height[0], height[len(height)-1]
    res := 0
    for l < r {
        if mL < mR {
            l++ 
            mL = max(mL, height[l])
            res += mL - height[l]
        } else {
            r--
            mR = max(mR, height[r])
            res += mR - height[r]
        }
    }

    return res
}