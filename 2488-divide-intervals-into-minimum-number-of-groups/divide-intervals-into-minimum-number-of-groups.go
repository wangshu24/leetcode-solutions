func minGroups(intervals [][]int) int {
    if len(intervals) == 1 {return 1}
    //Start by initialize 2 anchor int at 2 end
    //With the beginning anchor has largest int const to begin count down
    //And vice versa, smallest int const to for ending anchor to begin count up
    start := math.MaxInt
    end := math.MinInt
    
    for _,interval:= range intervals{
        start = min(start, interval[0])
        end = max(end, interval[1])
    }

    //We make a array to record occurence of beginning and end of each int
    //With positive value = more interval starts with this int
    //With negative value = more interval ends with this int
    //We add 2 to slices len since we take the end_int+1 in both val and index, which makes total_len = maxIndex+1 
    endArr := make([]int, end+2)
    for _,interval:= range intervals{
        //For every start val, we record its occurence
        endArr[interval[0]]++
        //For every end val, we record its closest subsequent start int with negative val
        //Denote the decrease in new group because end_val+1 is closest to end_val
        //And if that int does exist, which means it is included in the same group, does not need to record new occurence
        endArr[interval[1]+1]--
    }

    groups := 0
    maxConcurrentIntervals := 0
    for i:=start; i<=end;i++{
        groups += endArr[i]
        //Everytime, we record the max occurence possible because as we loop through, 
        //it will start encountering negative index, signify the closing of an interval group
        //therefore, need to record val before it goes to 0
        //that max is the min number of interval group
        //the algo already calc the most optimal/least no of group
        //hence, algo max = question expected min val
        maxConcurrentIntervals = max(maxConcurrentIntervals, groups)
    }

 
    return maxConcurrentIntervals
}
