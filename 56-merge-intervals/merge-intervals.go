func merge(intervals [][]int) [][]int {
    if len(intervals) == 1 {return intervals}
    
    slices.SortFunc(intervals, func(a,b []int) int {
        if n := cmp.Compare(a[0], b[0]); n != 0 {
            return n
        }
        return cmp.Compare(a[1], b[1])
    })

    start, end := intervals[0][0], intervals[0][1]
    res := [][]int{}
    for i:=1; i < len(intervals); i++ {

        if intervals[i][0] <= end {
            if intervals[i][1] > end {end = intervals[i][1]}
        } else {
            res = append(res, []int{start, end})
            start = intervals[i][0]
            end = intervals[i][1]
        }
    }

    res = append(res, []int{start, end})

    return res
}  