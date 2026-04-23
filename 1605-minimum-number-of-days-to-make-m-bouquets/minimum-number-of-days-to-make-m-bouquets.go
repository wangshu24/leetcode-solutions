func minDays(bloomday []int, m int, k int) int {
    if m * k > len(bloomday)  {return -1}
    if m * k == len(bloomday) {return slices.Max(bloomday)}

    minday, maxday := slices.Min(bloomday), slices.Max(bloomday)

    for minday < maxday {
        mid := minday + (maxday - minday)/2
        count := 0
        bouq := 0
        for _, day := range bloomday {
            if day <= mid {count++} else {count = 0}
            if count == k {
                count = 0
                bouq++    
            }
        }
        if bouq >= m {
            maxday = mid
        } else {
            minday = mid+1
        }
    }

    return minday
}