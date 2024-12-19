func maxChunksToSorted(ar []int) int {
    res := 0
    currMax := -1

    for key, val := range ar {
        currMax = max(val, currMax)
        if currMax == key {
            res++
        }
    }

    return res
}


func max (a,b int) int {
    if a > b {return a}
     return b
}