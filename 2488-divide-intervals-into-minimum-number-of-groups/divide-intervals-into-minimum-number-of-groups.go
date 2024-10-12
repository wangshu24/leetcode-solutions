func minGroups(intervals [][]int) int {
    if len(intervals) == 1 {return 1}
    start := math.MaxInt
    end := math.MinInt
    
    for _,interval:= range intervals{
        start = min(start, interval[0])
        end = max(end, interval[1])
    }

    // for i:=0; i <len(intervals);i++{
    //     if intervals[i][0] < start { start = intervals[i][0] }
    //     if intervals[i][1] > end { end = intervals[i][1] } 
    // }

    endArr := make([]int, end+2)
    for _,interval:= range intervals{
        endArr[interval[0]]++
        endArr[interval[1]+1]--
    }

    groups := 0
    maxConcurrentIntervals := 0
    for i:=start; i<=end;i++{
        groups += endArr[i]
        maxConcurrentIntervals = max(maxConcurrentIntervals, groups)
    }

 
    return maxConcurrentIntervals
}
