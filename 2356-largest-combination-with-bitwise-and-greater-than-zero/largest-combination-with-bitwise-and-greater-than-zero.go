func largestCombination(candidates []int) int {
    res := 0
    maxC := int(slices.Max(candidates))

    for bit := 1; bit <= maxC; bit = bit<<1 {
        count := 0
        for _, c := range candidates {
            if c & bit > 0 {
                count++
            }
        }
        if count > res {
            res = count
        }
    }
    return res
}